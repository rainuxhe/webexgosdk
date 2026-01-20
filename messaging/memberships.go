package messaging

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type Membership struct {
	ID                string    `json:"id,omitempty"`
	RoomID            string    `json:"roomId,omitempty"`
	PersonID          string    `json:"personId,omitempty"`
	PersonEmail       string    `json:"personEmail,omitempty"`
	PersonDisplayName string    `json:"personDisplayName,omitempty"`
	PersonOrgID       string    `json:"personOrgId,omitempty"`
	IsModerator       bool      `json:"isModerator,omitempty"`
	IsMonitor         bool      `json:"isMonitor,omitempty"`
	IsRoomHidden      bool      `json:"isRoomHidden,omitempty"`
	Created           time.Time `json:"created,omitempty"`
}

type MembershipsService struct {
	session *core.RestSession
}

func NewMembershipsService(session *core.RestSession) *MembershipsService {
	return &MembershipsService{
		session: session,
	}
}

type MembershipListOptions struct {
	RoomID      string
	PersonID    string
	PersonEmail string
	Max         int
}

func (s *MembershipsService) List(ctx context.Context, opts *MembershipListOptions) ([]*Membership, error) {
	params := url.Values{}
	if opts != nil {
		if opts.RoomID != "" {
			params.Set("roomId", opts.RoomID)
		}

		if opts.PersonID != "" {
			params.Set("personId", opts.PersonID)
		}

		if opts.PersonEmail != "" {
			params.Set("personEmail", opts.PersonEmail)
		}

		if opts.Max > 0 {
			params.Set("max", strconv.Itoa(opts.Max))
		}
	}

	var response struct {
		Items []*Membership `json:"items"`
	}

	if err := s.session.Get(ctx, "memberships", params, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

type MembershipCreateOptions struct {
	RoomID      string `json:"roomId"`
	PersonID    string `json:"personId,omitempty"`
	PersonEmail string `json:"personEmail,omitempty"`
	IsModerator bool   `json:"isModerator,omitempty"`
}

func (s *MembershipsService) Create(ctx context.Context, req *MembershipCreateOptions) (*Membership, error) {
	if req == nil || req.RoomID == "" {
		return nil, core.ErrInvalidParameter
	}

	if req.PersonID == "" && req.PersonEmail == "" {
		return nil, core.ErrInvalidParameter
	}

	var membership Membership
	if err := s.session.Post(ctx, "memberships", req, &membership); err != nil {
		return nil, err
	}

	return &membership, nil
}

func (s *MembershipsService) Get(ctx context.Context, membershipID string) (*Membership, error) {
	if membershipID == "" {
		return nil, core.ErrInvalidParameter
	}

	var membership Membership
	if err := s.session.Get(ctx, "memberships/"+membershipID, nil, &membership); err != nil {
		return nil, err
	}

	return &membership, nil
}

type MembershipUpdateRequest struct {
	IsModerator  bool `json:"isModerator,omitempty"`
	IsRoomHidden bool `json:"isRoomHidden,omitempty"`
}

func (s *MembershipsService) Update(ctx context.Context, membershipID string, req *MembershipUpdateRequest) (*Membership, error) {
	if membershipID == "" || req == nil {
		return nil, core.ErrInvalidParameter
	}

	var membership Membership
	if err := s.session.Put(ctx, "memberships/"+membershipID, req, &membership); err != nil {
		return nil, err
	}

	return &membership, nil
}

func (s *MembershipsService) Delete(ctx context.Context, membershipID string) error {
	if membershipID == "" {
		return core.ErrInvalidParameter
	}

	return s.session.Delete(ctx, "memberships/"+membershipID)
}

type TeamMembership struct {
	ID                string    `json:"id,omitempty"`
	TeamID            string    `json:"teamId,omitempty"`
	PersonID          string    `json:"personId,omitempty"`
	PersonEmail       string    `json:"personEmail,omitempty"`
	PersonDisplayName string    `json:"personDisplayName,omitempty"`
	PersonOrgID       string    `json:"personOrgId,omitempty"`
	IsModerator       bool      `json:"isModerator,omitempty"`
	Created           time.Time `json:"created,omitempty"`
}

type TeamMembershipsService struct {
	session *core.RestSession
}

func NewTeamMembershipsService(session *core.RestSession) *TeamMembershipsService {
	return &TeamMembershipsService{session: session}
}

type TeamMembershipListOptions struct {
	TeamID string
	Max    int
}

func (s *TeamMembershipsService) List(ctx context.Context, opts *TeamMembershipListOptions) ([]*TeamMembership, error) {
	if opts == nil || opts.TeamID == "" {
		return nil, core.ErrInvalidParameter
	}

	params := url.Values{}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}

	var response struct {
		Items []*TeamMembership `json:"items"`
	}

	if err := s.session.Get(ctx, "teams/memberships", params, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

type TeamMembershipCreateRequest struct {
	TeamID      string `json:"teamId"`
	PersonID    string `json:"personId,omitempty"`
	PersonEmail string `json:"personEmail,omitempty"`
	IsModerator bool   `json:"isModerator,omitempty"`
}

func (s *TeamMembershipsService) Create(ctx context.Context, req *TeamMembershipCreateRequest) (*TeamMembership, error) {
	if req == nil || req.TeamID == "" {
		return nil, core.ErrInvalidParameter
	}

	if req.PersonID == "" && req.PersonEmail == "" {
		return nil, core.ErrInvalidParameter
	}

	var membership TeamMembership
	if err := s.session.Post(ctx, "team/memberships", req, &membership); err != nil {
		return nil, err
	}

	return &membership, nil
}

func (s *TeamMembershipsService) Get(ctx context.Context, membershipID string) (*TeamMembership, error) {
	if membershipID == "" {
		return nil, core.ErrInvalidParameter
	}

	var membership TeamMembership
	if err := s.session.Get(ctx, "team/memberships/"+membershipID, nil, &membership); err != nil {
		return nil, err
	}

	return &membership, nil
}

type TeamMembershipUpdateRequest struct {
	IsModerator bool `json:"isModerator,omitempty"`
}

func (s *TeamMembershipsService) Update(ctx context.Context, membershipID string, req *TeamMembershipUpdateRequest) (*TeamMembership, error) {
	if membershipID == "" || req == nil {
		return nil, core.ErrInvalidParameter
	}

	var membership TeamMembership
	if err := s.session.Put(ctx, "team/memberships/"+membershipID, req, &membership); err != nil {
		return nil, err
	}

	return &membership, nil
}

func (s *TeamMembershipsService) Delete(ctx context.Context, membershipID string) error {
	if membershipID == "" {
		return core.ErrInvalidParameter
	}

	return s.session.Delete(ctx, "team/memberships/"+membershipID)
}
