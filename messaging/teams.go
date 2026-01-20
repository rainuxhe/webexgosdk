package messaging

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type Team struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	CreatorID   string    `json:"creatorId,omitempty"`
}

type TeamsService struct {
	session *core.RestSession
}

func NewTeamsService(session *core.RestSession) *TeamsService {
	return &TeamsService{
		session: session,
	}
}

type TeamListOptions struct {
	Max int
}

func (s *TeamsService) List(ctx context.Context, opts *TeamListOptions) ([]*Team, error) {
	params := url.Values{}
	if opts != nil {
		if opts.Max > 0 {
			params.Set("max", strconv.Itoa(opts.Max))
		}
	}

	var response struct {
		Items []*Team `json:"items"`
	}

	if err := s.session.Get(ctx, "teams", params, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

type TeamCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (s *TeamsService) Create(ctx context.Context, req *TeamCreateRequest) (*Team, error) {
	if req == nil || req.Name == "" {
		return nil, core.ErrInvalidParameter
	}

	var team Team
	if err := s.session.Post(ctx, "teams", req, &team); err != nil {
		return nil, err
	}
	return &team, nil
}

func (s *TeamsService) Get(ctx context.Context, teamID string) (*Team, error) {
	if teamID == "" {
		return nil, core.ErrInvalidParameter
	}

	var team Team
	if err := s.session.Get(ctx, "teams/"+teamID, nil, &team); err != nil {
		return nil, err
	}
	return &team, nil
}

type TeamUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (s *TeamsService) Update(ctx context.Context, teamID string, req *TeamUpdateRequest) (*Team, error) {
	if teamID == "" || req == nil || req.Name == "" {
		return nil, core.ErrInvalidParameter
	}

	var team Team
	if err := s.session.Put(ctx, "teams/"+teamID, req, &team); err != nil {
		return nil, err
	}
	return &team, nil
}

func (s *TeamsService) Delete(ctx context.Context, teamID string) error {
	if teamID == "" {
		return core.ErrInvalidParameter
	}

	return s.session.Delete(ctx, "teams/"+teamID)
}
