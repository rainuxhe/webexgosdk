package messaging

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type Message struct {
	ID              string       `json:"id,omitempty"`
	RoomID          string       `json:"roomId,omitempty"`
	RoomType        string       `json:"roomType,omitempty"`
	ParentID        string       `json:"parentId,omitempty"`
	ToPersonID      string       `json:"toPersonId,omitempty"`
	ToPersonEmail   string       `json:"toPersonEmail,omitempty"`
	Text            string       `json:"text,omitempty"`
	Markdown        string       `json:"markdown,omitempty"`
	Files           []string     `json:"files,omitempty"`
	PersonID        string       `json:"personId,omitempty"`
	PersonEmail     string       `json:"personEmail,omitempty"`
	MentionedPeople []string     `json:"mentionedPeople,omitempty"`
	MentionedGroups []string     `json:"mentionedGroups,omitempty"`
	Attachments     []Attachment `json:"attachments,omitempty"`
	Created         time.Time    `json:"created,omitempty"`
	Updated         time.Time    `json:"updated,omitempty"`
	IsVoiceClip     bool         `json:"isVoiceClip,omitempty"`
}

type Attachment struct {
	ContentType string `json:"contentType,omitempty"`
	Content     any    `json:"content,omitempty"`
}

type MessagesService struct {
	session *core.RestSession
}

func NewMessagesService(session *core.RestSession) *MessagesService {
	return &MessagesService{
		session: session,
	}
}

type MessageListOptions struct {
	// RoomID is required. List messages in a room.
	RoomID string

	// ParentID lists messages with this parent (for threaded messages).
	ParentID string

	// MentionedPeople lists messages where the caller is mentioned.
	// Use "me" or a specific personId.
	MentionedPeople string

	// Before lists messages sent before this date/time (ISO8601).
	Before string

	// BeforeMessage lists messages sent before this message ID.
	BeforeMessage string

	// Max limits the maximum number of messages returned.
	Max int
}

func (s *MessagesService) List(ctx context.Context, opts *MessageListOptions) ([]*Message, error) {
	if opts == nil || opts.RoomID == "" {
		return nil, core.ErrInvalidParameter
	}

	params := url.Values{}
	params.Set("roomId", opts.RoomID)

	if opts.ParentID != "" {
		params.Set("parentId", opts.ParentID)
	}

	if opts.MentionedPeople != "" {
		params.Set("mentionedPeople", opts.MentionedPeople)
	}

	if opts.Before != "" {
		params.Set("before", opts.Before)
	}

	if opts.BeforeMessage != "" {
		params.Set("beforeMessage", opts.BeforeMessage)
	}

	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}

	var response struct {
		Items []*Message `json:"items"`
	}

	if err := s.session.Get(ctx, "messages", params, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

type MessageDirectListOptions struct {
	// PersonID lists messages in a 1:1 room with this person.
	PersonID string

	// PersonEmail lists messages in a 1:1 room with this email.
	PersonEmail string

	// ParentID lists messages with this parent
	ParentID string
}

func (s *MessagesService) ListDirect(ctx context.Context, opts *MessageDirectListOptions) ([]*Message, error) {
	params := url.Values{}

	if opts != nil {
		if opts.PersonID != "" {
			params.Set("personId", opts.PersonID)
		}

		if opts.PersonEmail != "" {
			params.Set("personEmail", opts.PersonEmail)
		}

		if opts.ParentID != "" {
			params.Set("parentId", opts.ParentID)
		}
	}

	var response struct {
		Items []*Message `json:"items"`
	}

	if err := s.session.Get(ctx, "messages/direct", params, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

type MessageCreateRequest struct {
	RoomID        string       `json:"roomId,omitempty"`
	ToPersonID    string       `json:"toPersonId,omitempty"`
	ToPersonEmail string       `json:"toPersonEmail,omitempty"`
	Text          string       `json:"text,omitempty"`
	Markdown      string       `json:"markdown,omitempty"`
	Files         []string     `json:"files,omitempty"`
	Attachments   []Attachment `json:"attachments,omitempty"`
	ParentID      string       `json:"parentId,omitempty"`
}

func (s *MessagesService) Create(ctx context.Context, req *MessageCreateRequest) (*Message, error) {
	if req == nil {
		return nil, core.ErrInvalidParameter
	}

	if req.RoomID == "" && req.ToPersonID == "" && req.ToPersonEmail == "" {
		return nil, core.ErrInvalidParameter
	}

	if req.Text == "" && req.Markdown == "" && len(req.Files) == 0 && len(req.Attachments) == 0 {
		return nil, core.ErrInvalidParameter
	}

	var message Message
	if err := s.session.Post(ctx, "messages", req, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// Get returns details of a message by ID.
func (s *MessagesService) Get(ctx context.Context, messageID string) (*Message, error) {
	if messageID == "" {
		return nil, core.ErrInvalidParameter
	}

	var message Message
	if err := s.session.Get(ctx, "messages/"+messageID, nil, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

type MessageUpdateRequest struct {
	RoomID   string `json:"roomId,omitempty"`
	Text     string `json:"text,omitempty"`
	Markdown string `json:"markdown,omitempty"`
}

// Update edits an existing message.
func (s *MessagesService) Update(ctx context.Context, messageID string, req *MessageUpdateRequest) (*Message, error) {
	if messageID == "" {
		return nil, core.ErrInvalidParameter
	}

	var message Message
	if err := s.session.Put(ctx, "messages/"+messageID, req, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// Delete removes a message.
func (s *MessagesService) Delete(ctx context.Context, messageId string) error {
	if messageId == "" {
		return core.ErrInvalidParameter
	}

	return s.session.Delete(ctx, "messages/"+messageId)
}
