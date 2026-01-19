package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	DefaultBaseURL = "https://webexapis.com/v1/"
	DefaultTimeout = 60 * time.Second
)

type RestSession struct {
	client      *http.Client
	baseURL     string
	accessToken string
	userAgent   string
	mu          sync.RWMutex
}

type RestSessionConfig struct {
	AccessToken string
	BaseURL     string
	Timeout     time.Duration
	UserAgent   string
	HTTPClient  *http.Client
}

// NewRestSession creates a new RestSession with the provided configuration.
func NewRestSession(config *RestSessionConfig) *RestSession {
	baseURL := DefaultBaseURL
	if config.BaseURL != "" {
		baseURL = config.BaseURL
	}

	timeout := DefaultTimeout
	if config.Timeout > 0 {
		timeout = config.Timeout
	}

	client := config.HTTPClient
	if client == nil {
		client = &http.Client{
			Timeout: timeout,
		}
	}

	userAgent := "webexgosdk/1.0"
	if config.UserAgent != "" {
		userAgent = config.UserAgent
	}

	return &RestSession{
		client:      client,
		baseURL:     baseURL,
		accessToken: config.AccessToken,
		userAgent:   userAgent,
	}
}

func (s *RestSession) SetAccessToken(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.accessToken = token
}

func (s *RestSession) GetAccessToken() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.accessToken
}

func (s *RestSession) SetHTTPClient(client *http.Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.client = client
}

func (s *RestSession) doRequest(ctx context.Context, method, path string, params url.Values, body any, result any) error {
	fullURL := s.baseURL + path
	if params != nil {
		fullURL += "?" + params.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request")
	}

	s.mu.RLock()
	token := s.accessToken
	s.mu.RUnlock()

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.userAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return s.handleErrorResponse(resp)
	}

	if result != nil && resp.StatusCode != http.StatusNoContent {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if len(respBody) > 0 {
			if err := json.Unmarshal(respBody, result); err != nil {
				return fmt.Errorf("failed to unmarshal response body: %w", err)
			}
		}
	}

	return nil
}

// handleErrorResponse handles an error response from the API.
func (s *RestSession) handleErrorResponse(resp *http.Response) error {
	respBody, _ := io.ReadAll(resp.Body)
	trackingID := resp.Header.Get("Trackingid")

	var errorResp struct {
		Message string `json:"message"`
		Errors []ErrorDetail `json:"errors"`
	}

	if len(respBody) > 0 {
		json.Unmarshal(respBody, &errorResp)
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return NewRateLimitError(resp, errorResp.Message, trackingID)
	}
	return NewAPIError(resp, errorResp.Message, trackingID, errorResp.Errors)
}

func (s *RestSession) Get(ctx context.Context, path string, params url.Values, result any) error {
	return s.doRequest(ctx, http.MethodGet, path, params, nil, result)
}

func (s *RestSession) Post(ctx context.Context, path string, body any, result any) error {
	return s.doRequest(ctx, http.MethodPost, path, nil, body, result)
}

func (s *RestSession) Put(ctx context.Context, path string, body any, result any) error {
	return s.doRequest(ctx, http.MethodPut, path, nil, body, result)
}

func (s *RestSession) Delete(ctx context.Context, path string) error {
	return s.doRequest(ctx, http.MethodDelete, path, nil, nil, nil)
}

func (s *RestSession) PostMultipart(ctx context.Context, path string, fields map[string]string, files map[string]string, result any) error {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for key, value := range fields {
		if err := writer.WriteField(key, value); err != nil {
			return fmt.Errorf("failed to write field %s: %w", key, err)
		}
	}

	for fieldName, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", filePath, err)
		}
		defer file.Close()

		part, err := writer.CreateFormFile(fieldName, filepath.Base(filePath))
		if err != nil {
			return fmt.Errorf("failed to create form file: %w", err)
		}

		if _, err := io.Copy(part, file); err != nil {
			return fmt.Errorf("failed to copy file content: %w", err)
		}

	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	fullURL := s.baseURL + path
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, &buf)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	s.mu.RLock()
	token := s.accessToken
	s.mu.RUnlock()

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.userAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return s.handleErrorResponse(resp)
	}

	if result != nil && resp.StatusCode != http.StatusNoContent {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if len(respBody) > 0 {
			if err := json.Unmarshal(respBody, result); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}
		}
	}

	return nil
}