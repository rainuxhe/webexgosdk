package core

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	ErrInvalidParameter = errors.New("invalid or missing required parameter")
)

type APIError struct {
	StatusCode int
	Message    string
	TrackingID string
	Errors     []ErrorDetail
}

type ErrorDetail struct {
	Description string `json:"description"`
}

func (e *APIError) Error() string {
	msg := fmt.Sprintf("webex api error: status=%d", e.StatusCode)
	if e.Message != "" {
		msg += fmt.Sprintf(", message=%s", e.Message)
	}

	if e.TrackingID != "" {
		msg += fmt.Sprintf(", trackingId=%s", e.TrackingID)
	}

	return msg
}

type RateLimitError struct {
	*APIError
	RetryAfter time.Duration
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limit exceeded: retry after %v", e.RetryAfter)
}

func NewAPIError(resp *http.Response, message string, trackingID string, errors []ErrorDetail) *APIError {
	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    message,
		TrackingID: trackingID,
		Errors:     errors,
	}
}

func NewRateLimitError(resp *http.Response, message string, trackingID string) *RateLimitError {
	retryAfter := 60 * time.Second
	if retryHeader := resp.Header.Get("Retry-After"); retryHeader != "" {
		if seconds, err := strconv.Atoi(retryHeader); err == nil {
			retryAfter = time.Duration(seconds) * time.Second
		}
	}

	return &RateLimitError{
		APIError: &APIError{
			StatusCode: resp.StatusCode,
			Message:    message,
			TrackingID: trackingID,
		},
		RetryAfter: retryAfter,
	}
}
