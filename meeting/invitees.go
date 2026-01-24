package meeting

import (
	"context"
	"net/url"
	"strconv"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type MeetingInvitee struct {
	ID          string `json:"id,omitempty"`
	MeetingID   string `json:"meetingId,omitempty"`
	Email       string `json:"email,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	CoHost      bool   `json:"coHost,omitempty"`
	Panelist    bool   `json:"panelist,omitempty"`
	SendEmail   bool   `json:"sendEmail,omitempty"`
}

type MeetingInviteesService struct {
	session *core.RestSession
}

func NewMeetingInviteesService(session *core.RestSession) *MeetingInviteesService {
	return &MeetingInviteesService{
		session: session,
	}
}

type InviteeListOptions struct {
	MeetingID    string
	HostEmail    string
	PanelistOnly bool
	Max          int
}

func (s *MeetingInviteesService) List(ctx context.Context, opts *InviteeListOptions) ([]*MeetingInvitee, error) {
	if opts == nil || opts.MeetingID == "" {
		return nil, core.ErrInvalidParameter
	}

	params := url.Values{}
	params.Set("meetingId", opts.MeetingID)

	if opts.HostEmail != "" {
		params.Set("hostEmail", opts.HostEmail)
	}
	if opts.PanelistOnly {
		params.Set("panelistOnly", "true")
	}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}

	var response struct {
		Items []*MeetingInvitee `json:"items"`
	}

	if err := s.session.Get(ctx, "meetingInvitees", params, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

type InviteeCreateRequest struct {
	MeetingID   string `json:"meetingId"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName,omitempty"`
	CoHost      bool   `json:"coHost,omitempty"`
	Panelist    bool   `json:"panelist,omitempty"`
	SendEmail   bool   `json:"sendEmail,omitempty"`
	HostEmail   string `json:"hostEmail,omitempty"`
}

func (s *MeetingInviteesService) Create(ctx context.Context, req *InviteeCreateRequest) (*MeetingInvitee, error) {
	if req == nil || req.MeetingID == "" || req.Email == "" {
		return nil, core.ErrInvalidParameter
	}

	var invitee MeetingInvitee
	if err := s.session.Post(ctx, "meetingInvitees", req, &invitee); err != nil {
		return nil, err
	}

	return &invitee, nil
}

type BulkCreateRequest struct {
	MeetingID string            `json:"meetingId"`
	HostEmail string            `json:"hostEmail,omitempty"`
	SendEmail bool              `json:"sendEmail,omitempty"`
	Items     []BulkInviteeItem `json:"items"`
}

type BulkInviteeItem struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName,omitempty"`
	CoHost      bool   `json:"coHost,omitempty"`
	Panelist    bool   `json:"panelist,omitempty"`
}

// BulkCreate creates multiple meeting invitees at once.
func (s *MeetingInviteesService) BulkCreate(ctx context.Context, req *BulkCreateRequest) ([]*MeetingInvitee, error) {
	if req == nil || req.MeetingID == "" || len(req.Items) == 0 {
		return nil, core.ErrInvalidParameter
	}

	var response struct {
		Items []*MeetingInvitee `json:"items"`
	}

	if err := s.session.Post(ctx, "meetingInvitees", req, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

func (s *MeetingInviteesService) Get(ctx context.Context, inviteeID string) (*MeetingInvitee, error) {
	if inviteeID == "" {
		return nil, core.ErrInvalidParameter
	}

	var invitee MeetingInvitee
	if err := s.session.Get(ctx, "meetingInvitees/"+inviteeID, nil, &invitee); err != nil {
		return nil, err
	}

	return &invitee, nil
}

type InviteeUpdateRequest struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName,omitempty"`
	CoHost      bool   `json:"coHost,omitempty"`
	Panelist    bool   `json:"panelist,omitempty"`
	SendEmail   bool   `json:"sendEmail,omitempty"`
	HostEmail   string `json:"hostEmail,omitempty"`
}

func (s *MeetingInviteesService) Update(ctx context.Context, inviteeID string, req *InviteeUpdateRequest) (*MeetingInvitee, error) {
	if inviteeID == "" || req == nil {
		return nil, core.ErrInvalidParameter
	}

	var invitee MeetingInvitee
	if err := s.session.Put(ctx, "meetingInvitees/"+inviteeID, req, &invitee); err != nil {
		return nil, err
	}

	return &invitee, nil
}

func (s *MeetingInviteesService) Delete(ctx context.Context, inviteeID string) error {
	if inviteeID == "" {
		return core.ErrInvalidParameter
	}

	if err := s.session.Delete(ctx, "meetingInvitees/"+inviteeID); err != nil {
		return err
	}

	return nil
}
