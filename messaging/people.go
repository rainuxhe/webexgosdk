package messaging

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type Person struct {
	ID            string        `json:"id,omitempty"`
	Emails        []string      `json:"emails,omitempty"`
	PhoneNumbers  []PhoneNumber `json:"phoneNumbers,omitempty"`
	SIPAddresses  []SIPAddress  `json:"sipAddresses,omitempty"`
	DisplayName   string        `json:"displayName,omitempty"`
	NickName      string        `json:"nickName,omitempty"`
	FirstName     string        `json:"firstName,omitempty"`
	LastName      string        `json:"lastName,omitempty"`
	Avatar        string        `json:"avatar,omitempty"`
	OrgID         string        `json:"orgId,omitempty"`
	Roles         []string      `json:"roles,omitempty"`
	Licenses      []string      `json:"licenses,omitempty"`
	Department    string        `json:"department,omitempty"`
	Manager       string        `json:"manager,omitempty"`
	ManagerID     string        `json:"managerId,omitempty"`
	Title         string        `json:"title,omitempty"`
	Addresses     []Address     `json:"addresses,omitempty"`
	Created       time.Time     `json:"created,omitempty"`
	LastModified  time.Time     `json:"lastModified,omitempty"`
	LastActivity  time.Time     `json:"lastActivity,omitempty"`
	Status        string        `json:"status,omitempty"`
	InvitePending bool          `json:"invitePending,omitempty"`
	LoginEnabled  bool          `json:"loginEnabled,omitempty"`
	Type          string        `json:"type,omitempty"`
	Timezone      string        `json:"timezone,omitempty"`
}

type PhoneNumber struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type SIPAddress struct {
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type Address struct {
	Type       string `json:"type,omitempty"`
	Country    string `json:"country,omitempty"`
	Locality   string `json:"locality,omitempty"`
	Region     string `json:"region,omitempty"`
	StreetAddr string `json:"streetAddress,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
}

type PeopleService struct {
	session *core.RestSession
}

func NewPeopleService(session *core.RestSession) *PeopleService {
	return &PeopleService{
		session: session,
	}
}

type PeopleListOptions struct {
	Email       string
	DisplayName string
	ID          string
	OrgID       string
	CallingData bool
	LocationID  string
	Max         int
}

func (s *PeopleService) List(ctx context.Context, opts *PeopleListOptions) ([]*Person, error) {
	params := url.Values{}

	if opts != nil {
		if opts.Email != "" {
			params.Set("email", opts.Email)
		}
		if opts.DisplayName != "" {
			params.Set("displayName", opts.DisplayName)
		}
		if opts.ID != "" {
			params.Set("id", opts.ID)
		}
		if opts.OrgID != "" {
			params.Set("orgId", opts.OrgID)
		}
		if opts.CallingData {
			params.Set("callingData", "true")
		}
		if opts.LocationID != "" {
			params.Set("locationId", opts.LocationID)
		}
		if opts.Max > 0 {
			params.Set("max", strconv.Itoa(opts.Max))
		}
	}

	var response struct {
		Items []*Person `json:"items"`
	}

	if err := s.session.Get(ctx, "people", params, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (s *PeopleService) Get(ctx context.Context, personID string) (*Person, error) {
	if personID == "" {
		return nil, core.ErrInvalidParameter
	}

	var person Person
	if err := s.session.Get(ctx, "people/"+personID, nil, &person); err != nil {
		return nil, err
	}
	return &person, nil
}

func (s *PeopleService) GetMe(ctx context.Context) (*Person, error) {
	var person Person
	if err := s.session.Get(ctx, "people/me", nil, &person); err != nil {
		return nil, err
	}
	return &person, nil
}

type PersonCreateRequest struct {
	Emails       []string      `json:"emails"`
	DisplayName  string        `json:"displayName,omitempty"`
	FirstName    string        `json:"firstName,omitempty"`
	LastName     string        `json:"lastName,omitempty"`
	Avatar       string        `json:"avatar,omitempty"`
	OrgID        string        `json:"orgId,omitempty"`
	Roles        []string      `json:"roles,omitempty"`
	Licenses     []string      `json:"licenses,omitempty"`
	Department   string        `json:"department,omitempty"`
	Manager      string        `json:"manager,omitempty"`
	ManagerID    string        `json:"managerId,omitempty"`
	Title        string        `json:"title,omitempty"`
	Addresses    []Address     `json:"addresses,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phoneNumbers,omitempty"`
	SiteUrls     []string      `json:"siteUrls,omitempty"`
	LocationID   string        `json:"locationId,omitempty"`
}

func (s *PeopleService) Create(ctx context.Context, req *PersonCreateRequest) (*Person, error) {
	if req == nil || len(req.Emails) == 0 {
		return nil, core.ErrInvalidParameter
	}
	var person Person
	if err := s.session.Post(ctx, "people", req, &person); err != nil {
		return nil, err
	}
	return &person, nil
}

type PersonUpdateRequest struct {
	Emails       []string      `json:"emails,omitempty"`
	DisplayName  string        `json:"displayName,omitempty"`
	FirstName    string        `json:"firstName,omitempty"`
	LastName     string        `json:"lastName,omitempty"`
	Avatar       string        `json:"avatar,omitempty"`
	OrgID        string        `json:"orgId,omitempty"`
	Roles        []string      `json:"roles,omitempty"`
	Licenses     []string      `json:"licenses,omitempty"`
	Department   string        `json:"department,omitempty"`
	Manager      string        `json:"manager,omitempty"`
	ManagerID    string        `json:"managerId,omitempty"`
	Title        string        `json:"title,omitempty"`
	Addresses    []Address     `json:"addresses,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phoneNumbers,omitempty"`
	SiteUrls     []string      `json:"siteUrls,omitempty"`
	LocationID   string        `json:"locationId,omitempty"`
	LoginEnabled *bool         `json:"loginEnabled,omitempty"`
}

func (s *PeopleService) Update(ctx context.Context, personID string, req *PersonUpdateRequest) (*Person, error) {
	if personID == "" || req == nil {
		return nil, core.ErrInvalidParameter
	}

	var person Person
	if err := s.session.Put(ctx, "people/"+personID, req, &person); err != nil {
		return nil, err
	}
	return &person, nil
}

func (s *PeopleService) Delete(ctx context.Context, personID string) error {
	if personID == "" {
		return core.ErrInvalidParameter
	}
	return s.session.Delete(ctx, "people/"+personID)
}
