package webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/config"
)

type webhookOpt struct {
	Filter  string `json:"filter,omitempty"`
	Secret  string `json:"secret,omitempty"`
	OwnedBy string `json:"ownedBy,omitempty"`
	Max     int    `json:"max,omitempty"`
}

type option func(*webhookOpt)

func WithFilter(filter string) option {
	return func(opt *webhookOpt) {
		opt.Filter = filter
	}
}

func WithSecret(secret string) option {
	return func(opt *webhookOpt) {
		opt.Secret = secret
	}
}

func WithOwnedBy(ownedBy string) option {
	return func(opt *webhookOpt) {
		opt.OwnedBy = ownedBy
	}
}

func WithMax(max int) option {
	return func(opt *webhookOpt) {
		opt.Max = max
	}
}

type webhooks struct {
	config *config.Config
}

func NewWebhooks() *webhooks {
	return &webhooks{
		config: config.NewConfig(),
	}
}

func (w *webhooks) catchErr(respBody io.Reader) {
	var resp map[string]any
	err := json.NewDecoder(respBody).Decode(&resp)
	if err != nil {
		fmt.Println("JSON Decode error: ", err)
	} else {
		fmt.Printf("Invoking the API failed with error: %+v\n", resp)
	}
}

func (w *webhooks) Create(event, name, resource, targetUrl string, opts ...option) (response Response, err error) {
	if !w.isValidateCreate(event, resource) {
		return response, fmt.Errorf("invalid event or resource")
	}

	url := fmt.Sprintf("%s/v1/webhooks", w.config.BaseURL())

	otherOpt := &webhookOpt{}
	for _, opt := range opts {
		opt(otherOpt)
	}

	reqBody := struct {
		Event     string `json:"event"`
		Name      string `json:"name"`
		Resource  string `json:"resource"`
		TargetUrl string `json:"targetUrl"`
		webhookOpt
	}{
		Event:      event,
		Name:       name,
		Resource:   resource,
		TargetUrl:  targetUrl,
		webhookOpt: *otherOpt,
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", w.config.AccessToken()))

	client := &http.Client{
		Timeout: time.Duration(w.config.TimeOut()) * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		w.catchErr(resp.Body)
		return
	}

	return
}

func (w *webhooks) isValidateCreate(event, resource string) bool {
	events := []string{
		"created",
		"updated",
		"deleted",
		"started",
		"ended",
		"joined",
		"left",
		"migrated",
		"authorized",
		"deauthorized",
		"statusChanged",
	}

	resources := []string{
		"attachmentActions",
		"dataSources",
		"memberships",
		"messages",
		"rooms",
		"meetings",
		"recordings",
		"convergedRecordings",
		"meetingParticipants",
		"meetingTranscripts",
		"uc_counters",
		"serviceApp",
		"adminBatchJobs",
	}
	return slices.Contains(events, event) && slices.Contains(resources, resource)
}

func (w *webhooks) Delete() {}

func (w *webhooks) Get() {}

// List all of your webhooks.
// Parameters:
//   - WithMax, Limit the maximum number of webhooks in the response.
func (w *webhooks) List(opts ...option) (response []Response, err error) {
	targetUrl := fmt.Sprintf("%s/v1/webhooks", w.config.BaseURL())
	reqbody := &webhookOpt{
		Max: 100,
	}
	for _, opt := range opts {
		opt(reqbody)
	}
	formData := url.Values{}
	formData.Set("ownedBy", "org")
	formData.Set("max", fmt.Sprintf("%d", reqbody.Max))

	req, err := http.NewRequest(http.MethodGet, targetUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", w.config.AccessToken()))

	client := &http.Client{
		Timeout: time.Duration(w.config.TimeOut()) * time.Second,
	}
	client.Do(req)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.catchErr(resp.Body)
		return
	}

	var respBody map[string][]Response
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return
	}
	return respBody["items"], nil
}

func (w *webhooks) Update() {}
