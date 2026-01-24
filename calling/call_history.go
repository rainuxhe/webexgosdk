package calling

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type CallHistoryRecord struct {
	ID              string      `json:"id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Number          string      `json:"number,omitempty"`
	Type            string      `json:"type,omitempty"`
	Direction       string      `json:"direction,omitempty"`
	Duration        int         `json:"duration,omitempty"`
	StartTime       time.Time   `json:"startTime,omitempty"`
	AnswerTime      time.Time   `json:"answerTime,omitempty"`
	EndTime         time.Time   `json:"endTime,omitempty"`
	IsCallback      bool        `json:"isCallback,omitempty"`
	CallbackNumber  string      `json:"callbackNumber,omitempty"`
	CallSessionID   string      `json:"callSessionId,omitempty"`
	LocalCallID     string      `json:"localCallId,omitempty"`
	RemoteCallID    string      `json:"remoteCallId,omitempty"`
	UserType        string      `json:"userType,omitempty"`
	UserID          string      `json:"userId,omitempty"`
	OrgID           string      `json:"orgId,omitempty"`
	IsAnswered      bool        `json:"isAnswered,omitempty"`
	IsInternational bool        `json:"isInternational,omitempty"`
	OriginalReason  *CallReason `json:"originalReason,omitempty"`
	RedirectReason  *CallReason `json:"redirectReason,omitempty"`
	ReleasedParty   string      `json:"releasedParty,omitempty"`
}

type CallReason struct {
	Reason string `json:"reason,omitempty"`
	Code   string `json:"code,omitempty"`
}

type CallHistoryService struct {
	session *core.RestSession
}

func NewCallHistoryService(session *core.RestSession) *CallHistoryService {
	return &CallHistoryService{
		session: session,
	}
}

type CallHistoryListOptions struct {
	Type string
	Max  int
}

func (s *CallHistoryService) List(ctx context.Context, opts *CallHistoryListOptions) ([]*CallHistoryRecord, error) {
	params := url.Values{}
	if opts != nil {
		if opts.Type != "" {
			params.Set("type", opts.Type)
		}
		if opts.Max > 0 {
			params.Set("max", strconv.Itoa(opts.Max))
		}
	}

	var response struct {
		Items []*CallHistoryRecord `json:"items"`
	}

	if err := s.session.Get(ctx, "telephony/calls/history", params, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

func (s *CallHistoryService) ListPlacedCalls(ctx context.Context, max int) ([]*CallHistoryRecord, error) {
	opts := &CallHistoryListOptions{
		Type: "placed",
	}

	if max > 0 {
		opts.Max = max
	}

	return s.List(ctx, opts)
}

func (s *CallHistoryService) ListMissedCalls(ctx context.Context, max int) ([]*CallHistoryRecord, error) {
	opts := &CallHistoryListOptions{
		Type: "missed",
	}

	if max > 0 {
		opts.Max = max
	}

	return s.List(ctx, opts)
}

func (s *CallHistoryService) ListReceivedCalls(ctx context.Context, max int) ([]*CallHistoryRecord, error) {
	opts := &CallHistoryListOptions{
		Type: "received",
	}

	if max > 0 {
		opts.Max = max
	}

	return s.List(ctx, opts)
}

func (s *CallHistoryService) DeleteAllHistory(ctx context.Context) error {
	return s.session.Delete(ctx, "telephony/calls/history")
}
