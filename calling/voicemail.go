package calling

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type VoiceMessage struct {
	ID           string             `json:"id,omitempty"`
	Duration     int                `json:"duration,omitempty"`
	CallingParty *VoiceMessageParty `json:"callingParty,omitempty"`
	Urgent       bool               `json:"urgent,omitempty"`
	Confidential bool               `json:"confidential,omitempty"`
	Read         bool               `json:"read,omitempty"`
	Created      time.Time          `json:"created,omitempty"`
	MediaType    string             `json:"mediaType,omitempty"`
	MessageType  string             `json:"messageType,omitempty"`
}

type VoiceMessageParty struct {
	Name           string `json:"name,omitempty"`
	Number         string `json:"number,omitempty"`
	PersonID       string `json:"personId,omitempty"`
	PlaceID        string `json:"placeId,omitempty"`
	PrivacyEnabled bool   `json:"privacyEnabled,omitempty"`
}

type VoiceMessageSummary struct {
	NewMessages       int `json:"newMessages,omitempty"`
	OldMessages       int `json:"oldMessages,omitempty"`
	NewUrgentMessages int `json:"newUrgentMessages,omitempty"`
	OldUrgentMessages int `json:"oldUrgentMessages,omitempty"`
}

type VoicemailService struct {
	session *core.RestSession
}

func NewVoicemailService(session *core.RestSession) *VoicemailService {
	return &VoicemailService{
		session: session,
	}
}

type VoicemailListOptions struct {
	Max int
}

func (s *VoicemailService) List(ctx context.Context, opts *VoicemailListOptions) ([]*VoiceMessage, error) {
	params := url.Values{}

	if opts != nil {
		if opts.Max > 0 {
			params.Set("max", strconv.Itoa(opts.Max))
		}
	}

	var response struct {
		Items []*VoiceMessage `json:"items"`
	}
	if err := s.session.Get(ctx, "telephony/voiceMessages", params, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (s *VoicemailService) GetSummary(ctx context.Context) (*VoiceMessageSummary, error) {
	var summary VoiceMessageSummary
	if err := s.session.Get(ctx, "telephony/voiceMessages/summary", nil, &summary); err != nil {
		return nil, err
	}
	return &summary, nil
}

func (s *VoicemailService) MarkAsRead(ctx context.Context, messageID string) error {
	if messageID == "" {
		return core.ErrInvalidParameter
	}

	req := map[string]bool{"read": true}
	return s.session.Put(ctx, "telephony/voiceMessages/"+messageID, req, nil)
}

func (s *VoicemailService) MarkAsUnread(ctx context.Context, messageID string) error {
	if messageID == "" {
		return core.ErrInvalidParameter
	}
	req := map[string]bool{"read": false}
	return s.session.Put(ctx, "telephony/voiceMessages/"+messageID, req, nil)
}

func (s *VoicemailService) Delete(ctx context.Context, messageID string) error {
	if messageID == "" {
		return core.ErrInvalidParameter
	}
	return s.session.Delete(ctx, "telephony/voiceMessages/"+messageID)
}
