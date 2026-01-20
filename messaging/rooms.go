package messaging

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type Room struct {
	ID                 string    `json:"id,omitempty"`
	Title              string    `json:"title,omitempty"`
	Type               string    `json:"type,omitempty"`
	IsLocked           bool      `json:"isLocked,omitempty"`
	IsPublic           bool      `json:"isPublic,omitempty"`
	IsAnnouncementOnly bool      `json:"isAnnouncementOnly,omitempty"`
	IsReadOnly         bool      `json:"isReadOnly,omitempty"`
	TeamID             string    `json:"teamId,omitempty"`
	ClassificationID   string    `json:"classificationId,omitempty"`
	Description        string    `json:"description,omitempty"`
	LastActivity       time.Time `json:"lastActivity,omitempty"`
	Created            time.Time `json:"created,omitempty"`
	CreatorID          string    `json:"creatorId,omitempty"`
	SipAddress         string    `json:"sipAddress,omitempty"`
	OwnerID            string    `json:"ownerId,omitempty"`
}

type RoomMeetingInfo struct {
	RoomID               string `json:"roomId,omitempty"`
	MeetingLink          string `json:"meetingLink,omitempty"`
	SipAddress           string `json:"sipAddress,omitempty"`
	MeetingNumber        string `json:"meetingNumber,omitempty"`
	CallInTollFreeNumber string `json:"callInTollFreeNumber,omitempty"`
	CallInTollNumber     string `json:"callInTollNumber,omitempty"`
}

type RoomsService struct {
	session *core.RestSession
}

func NewRoomsService(session *core.RestSession) *RoomsService {
	return &RoomsService{
		session: session,
	}
}

type RoomListOptions struct {
	TeamID string
	Type   string
	SortBy string
	Max    int
}

func (s *RoomsService) List(ctx context.Context, opts *RoomListOptions) ([]*Room, error) {
	params := url.Values{}
	if opts != nil {
		if opts.TeamID != "" {
			params.Set("teamId", opts.TeamID)
		}

		if opts.Type != "" {
			params.Set("type", opts.Type)
		}

		if opts.SortBy != "" {
			params.Set("sortBy", opts.SortBy)
		}

		if opts.Max > 0 {
			params.Set("max", strconv.Itoa(opts.Max))
		}
	}

	var response struct {
		Items []*Room `json:"items"`
	}

	if err := s.session.Get(ctx, "rooms", params, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

type RoomCreateRequest struct {
	Title              string `json:"title"`
	TeamID             string `json:"teamId,omitempty"`
	ClassificationID   string `json:"classificationId,omitempty"`
	IsLocked           bool   `json:"isLocked,omitempty"`
	IsPublic           bool   `json:"isPublic,omitempty"`
	Description        string `json:"description,omitempty"`
	IsAnnouncementOnly bool   `json:"isAnnouncementOnly,omitempty"`
}

// Create creates a new room.
func (s *RoomsService) Create(ctx context.Context, req *RoomCreateRequest) (*Room, error) {
	if req == nil || req.Title == "" {
		return nil, core.ErrInvalidParameter
	}

	var room Room
	if err := s.session.Post(ctx, "rooms", req, &room); err != nil {
		return nil, err
	}

	return &room, nil
}

func (s *RoomsService) Get(ctx context.Context, roomID string) (*Room, error) {
	if roomID == "" {
		return nil, core.ErrInvalidParameter
	}

	var room Room
	if err := s.session.Get(ctx, "rooms/"+roomID, nil, &room); err != nil {
		return nil, err
	}

	return &room, nil
}

func (s *RoomsService) GetMeetingInfo(ctx context.Context, roomID string) (*RoomMeetingInfo, error) {
	if roomID == "" {
		return nil, core.ErrInvalidParameter
	}

	var meetingInfo RoomMeetingInfo
	if err := s.session.Get(ctx, "rooms/"+roomID+"/meetingInfo", nil, &meetingInfo); err != nil {
		return nil, err
	}

	return &meetingInfo, nil
}

type RoomUpdateRequest struct {
	Title              string `json:"title"`
	TeamID             string `json:"teamId,omitempty"`
	ClassificationID   string `json:"classificationId,omitempty"`
	IsLocked           bool   `json:"isLocked,omitempty"`
	IsPublic           bool   `json:"isPublic,omitempty"`
	Description        string `json:"description,omitempty"`
	IsAnnouncementOnly bool   `json:"isAnnouncementOnly,omitempty"`
	IsReadOnly         bool   `json:"isReadOnly,omitempty"`
}

func (s *RoomsService) Update(ctx context.Context, roomID string, req *RoomUpdateRequest) (*Room, error) {
	if roomID == "" || req == nil || req.Title == "" {
		return nil, core.ErrInvalidParameter
	}

	var room Room
	if err := s.session.Put(ctx, "rooms/"+roomID, req, &room); err != nil {
		return nil, err
	}

	return &room, nil
}

func (s *RoomsService) Delete(ctx context.Context, roomID string) error {
	if roomID == "" {
		return core.ErrInvalidParameter
	}

	return s.session.Delete(ctx, "rooms/"+roomID)
}
