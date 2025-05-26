package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/YXHYW/webexgosdk/internal/config"
	"github.com/YXHYW/webexgosdk/models"
)

type messageReqBody struct {
	RoomId          string              `json:"roomId,omitempty"`
	ParentId        string              `json:"parentId,omitempty"`
	ToPersonId      string              `json:"toPersonId,omitempty"`
	ToPersonEmail   string              `json:"toPersonEmail,omitempty"`
	Text            string              `json:"text,omitempty"`
	Markdown        string              `json:"markdown,omitempty"`
	Files           []string            `json:"files,omitempty"`
	Attachments     []models.Attachment `json:"attachments,omitempty"`
	MentionedPeople []string            `json:"mentionedPeople,omitempty"`
	Before          string              `json:"before,omitempty"`
	BeforeMessage   string              `json:"beforeMessage,omitempty"`
	Max             int                 `json:"max,omitempty"`
}

type option func(*messageReqBody)

func WithRoomId(roomId string) option {
	return func(m *messageReqBody) {
		m.RoomId = roomId
	}
}

func WithParentId(parentId string) option {
	return func(m *messageReqBody) {
		m.ParentId = parentId
	}
}

func WithToPersonId(toPersonId string) option {
	return func(m *messageReqBody) {
		m.ToPersonId = toPersonId
	}
}

func WithToPersonEmail(toPersonEmail string) option {
	return func(m *messageReqBody) {
		m.ToPersonEmail = toPersonEmail
	}
}

func WithText(text string) option {
	return func(m *messageReqBody) {
		m.Text = text
	}
}

func WithMarkdown(markdown string) option {
	return func(m *messageReqBody) {
		m.Markdown = markdown
	}
}

func WithFiles(files []string) option {
	return func(m *messageReqBody) {
		m.Files = files
	}
}

func WithAttachments(attachments []models.Attachment) option {
	return func(m *messageReqBody) {
		m.Attachments = attachments
	}
}

func WithMentionedPeople(mentionedPeople []string) option {
	return func(m *messageReqBody) {
		m.MentionedPeople = mentionedPeople
	}
}

func WithBefore(before string) option {
	return func(m *messageReqBody) {
		m.Before = before
	}
}

func WithBeforeMessage(beforeMessage string) option {
	return func(m *messageReqBody) {
		m.BeforeMessage = beforeMessage
	}
}

func WithMax(max int) option {
	return func(m *messageReqBody) {
		m.Max = max
	}
}

func NewMessageRequestBody(opts ...option) *messageReqBody {
	reqBody := &messageReqBody{}
	for _, opt := range opts {
		opt(reqBody)
	}
	return reqBody
}

type messages struct {
	config *config.Config
}

func NewMessages() *messages {
	return &messages{
		config: config.NewConfig(),
	}
}

// Catch the response body if webexapi returns an error
//
// Parameters:
//   - respBody: The response body from the webexapi
func (m *messages) catchErr(respBody io.Reader) {
	var resp map[string]any
	err := json.NewDecoder(respBody).Decode(&resp)
	if err != nil {
		fmt.Println("JSON Decode error: ", err)
	} else {
		fmt.Printf("Invoking the API failed with error: %+v\n", resp)
	}
}

// Post a plain text or rich text message, and optionally, a file attachment attachment, to a room.
func (m *messages) Create(opts ...option) (response models.Response, err error) {
	url := fmt.Sprintf("%s/v1/messages", m.config.BaseURL())
	reqBody := &messageReqBody{}
	for _, opt := range opts {
		opt(reqBody)
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.config.AccessToken()))

	client := &http.Client{
		Timeout: time.Duration(m.config.TimeOut()) * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", resp.Status)
		m.catchErr(resp.Body)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	return
}

func (m *messages) Delete(messageId string) error {
	url := fmt.Sprintf("%s/v1/messages/%s", m.config.BaseURL(), messageId)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.config.AccessToken()))
	client := &http.Client{Timeout: time.Duration(m.config.TimeOut()) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		m.catchErr(resp.Body)
		return fmt.Errorf("failed to delete message with id %s", messageId)
	}

	return nil
}

func (m *messages) Edit(messageId string, opts ...option) (response models.Response, err error) {
	url := fmt.Sprintf("%s/v1/messages/%s", m.config.BaseURL(), messageId)
	reqBody := &messageReqBody{}
	for _, opt := range opts {
		opt(reqBody)
	}
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.config.AccessToken()))

	client := &http.Client{
		Timeout: time.Duration(m.config.TimeOut()) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		m.catchErr(resp.Body)
		return response, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	return
}

func (m *messages) Get(messageId string) (response models.Response, err error) {
	url := fmt.Sprintf("%s/v1/messages/%s", m.config.BaseURL(), messageId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.config.AccessToken()))
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: time.Duration(m.config.TimeOut()) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		m.catchErr(resp.Body)
		err = fmt.Errorf("messags get failed, status code: %d", resp.StatusCode)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}
	return
}

// List all messages in a 1:1 (direct) room. Use the personId or personEmail query parameter to specify the room. Each message will include content attachments if present.
// The list sorts the messages in descending order by creation date.
func (m *messages) ListDirect(opts ...option) (response []models.Response, err error) {
	url := fmt.Sprintf("%s/messages/direct", m.config.BaseURL())
	reqBody := &messageReqBody{}
	for _, opt := range opts {
		opt(reqBody)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.config.AccessToken()))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: time.Duration(m.config.TimeOut()) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		m.catchErr(resp.Body)
		return
	}

	var respBody map[string][]models.Response
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return
	}
	return respBody["items"], nil
}

// Lists all messages in a room. Each message will include content attachments if present.
func (m *messages) List(opts ...option) (response []models.Response, err error) {
	url := fmt.Sprintf("%s/messages", m.config.BaseURL())
	reqBody := &messageReqBody{}
	for _, opt := range opts {
		opt(reqBody)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.config.AccessToken()))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: time.Duration(m.config.TimeOut()) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		m.catchErr(resp.Body)
		return
	}

	var respBody map[string][]models.Response
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return
	}
	return respBody["items"], nil
}