package meeting

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type MeetingRegistrant struct {
	ID                  string               `json:"id,omitempty"`
	MeetingID           string               `json:"meetingId,omitempty"`
	RegisterTime        time.Time            `json:"registerTime,omitempty"`
	Status              string               `json:"status,omitempty"`
	FirstName           string               `json:"firstName,omitempty"`
	LastName            string               `json:"lastName,omitempty"`
	Email               string               `json:"email,omitempty"`
	JobTitle            string               `json:"jobTitle,omitempty"`
	CompanyName         string               `json:"companyName,omitempty"`
	Address1            string               `json:"address1,omitempty"`
	Address2            string               `json:"address2,omitempty"`
	City                string               `json:"city,omitempty"`
	State               string               `json:"state,omitempty"`
	ZipCode             string               `json:"zipCode,omitempty"`
	CountryRegion       string               `json:"countryRegion,omitempty"`
	WorkPhone           string               `json:"workPhone,omitempty"`
	Fax                 string               `json:"fax,omitempty"`
	SourceID            string               `json:"sourceId,omitempty"`
	SendEmail           bool                 `json:"sendEmail,omitempty"`
	CustomizedQuestions []CustomizedQuestion `json:"customizedQuestions,omitempty"`
}

type CustomizedQuestion struct {
	QuestionID string `json:"questionId,omitempty"`
	Answer     string `json:"answer,omitempty"`
}

type MeetingRegistrantsService struct {
	session *core.RestSession
}

func NewMeetingRegistrantsService(session *core.RestSession) *MeetingRegistrantsService {
	return &MeetingRegistrantsService{
		session: session,
	}
}

type RegistrantListOptions struct {
	MeetingID string
	HostEmail string
	OrderType string
	Orderby   string
	Current   bool
	Email     string
	Max       int
}

func (s *MeetingRegistrantsService) List(ctx context.Context, opts *RegistrantListOptions) ([]*MeetingRegistrant, error) {
	if opts == nil || opts.MeetingID == "" {
		return nil, core.ErrInvalidParameter
	}

	params := url.Values{}
	params.Set("meetingId", opts.MeetingID)

	if opts.HostEmail != "" {
		params.Set("hostEmail", opts.HostEmail)
	}

	if opts.OrderType != "" {
		params.Set("orderType", opts.OrderType)
	}

	if opts.Orderby != "" {
		params.Set("orderby", opts.Orderby)
	}

	if opts.Current {
		params.Set("current", "true")
	}

	if opts.Email != "" {
		params.Set("email", opts.Email)
	}

	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}

	var response struct {
		Items []*MeetingRegistrant `json:"items"`
	}

	if err := s.session.Get(ctx, "meetingRegistrants", params, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

type RegistrantCreateRequest struct {
	MeetingID           string               `json:"meetingId"`
	FirstName           string               `json:"firstName"`
	LastName            string               `json:"lastName"`
	Email               string               `json:"email"`
	JobTitle            string               `json:"jobTitle,omitempty"`
	CompanyName         string               `json:"companyName,omitempty"`
	Address1            string               `json:"address1,omitempty"`
	Address2            string               `json:"address2,omitempty"`
	City                string               `json:"city,omitempty"`
	State               string               `json:"state,omitempty"`
	ZipCode             string               `json:"zipCode,omitempty"`
	CountryRegion       string               `json:"countryRegion,omitempty"`
	WorkPhone           string               `json:"workPhone,omitempty"`
	Fax                 string               `json:"fax,omitempty"`
	SendEmail           bool                 `json:"sendEmail"`
	HostEmail           string               `json:"hostEmail,omitempty"`
	CustomizedQuestions []CustomizedQuestion `json:"customizedQuestions,omitempty"`
}

func (s *MeetingRegistrantsService) Create(ctx context.Context, req *RegistrantCreateRequest) (*MeetingRegistrant, error) {
	if req == nil || req.MeetingID == "" || req.FirstName == "" || req.LastName == "" || req.Email == "" {
		return nil, core.ErrInvalidParameter
	}

	var registrant MeetingRegistrant

	if err := s.session.Post(ctx, "meetingRegistrants", req, &registrant); err != nil {
		return nil, err
	}

	return &registrant, nil
}

type BulkRegistrantCreateRequest struct {
	MeetingID string               `json:"meetingId"`
	HostEmail string               `json:"hostEmail,omitempty"`
	SendEmail bool                 `json:"sendEmail,omitempty"`
	Items     []BulkRegistrantItem `json:"items"`
}

type BulkRegistrantItem struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	JobTitle    string `json:"jobTitle,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
}

func (s *MeetingRegistrantsService) BulkCreate(ctx context.Context, req *BulkRegistrantCreateRequest) ([]*MeetingRegistrant, error) {
	if req == nil || req.MeetingID == "" || len(req.Items) == 0 {
		return nil, core.ErrInvalidParameter
	}

	var response struct {
		Items []*MeetingRegistrant `json:"items"`
	}

	if err := s.session.Post(ctx, "meetingRegistrants", req, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

func (s *MeetingRegistrantsService) Get(ctx context.Context, registrantID string) (*MeetingRegistrant, error) {
	if registrantID == "" {
		return nil, core.ErrInvalidParameter
	}

	var registrant MeetingRegistrant

	if err := s.session.Get(ctx, "meetingRegistrants/"+registrantID, nil, &registrant); err != nil {
		return nil, err
	}

	return &registrant, nil
}

type RegistrantUpdateStatusRequest struct {
	Status    string `json:"status"`
	SendEmail bool   `json:"sendEmail,omitempty"`
	HostEmail string `json:"hostEmail,omitempty"`
}

func (s *MeetingRegistrantsService) UpdateStatus(ctx context.Context, registrantID string, req *RegistrantUpdateStatusRequest) (*MeetingRegistrant, error) {
	if registrantID == "" || req == nil || req.Status == "" {
		return nil, core.ErrInvalidParameter
	}

	var response MeetingRegistrant

	if err := s.session.Put(ctx, "meetingRegistrants/"+registrantID, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *MeetingRegistrantsService) Delete(ctx context.Context, registrantID string) error {
	if registrantID == "" {
		return core.ErrInvalidParameter
	}

	return s.session.Delete(ctx, "meetingRegistrants/"+registrantID)
}

type QueryRegistrationFormOption struct {
	HostEmail string
	Current   bool
}

type RegistrationForm struct {
	AutoRegisterEnabled  bool                   `json:"autoRegisterEnabled,omitempty"`
	AutoAcceptRequest    bool                   `json:"autoAcceptRequest,omitempty"`
	RequireFirstName     bool                   `json:"requireFirstName,omitempty"`
	RequireLastName      bool                   `json:"requireLastName,omitempty"`
	RequireEmail         bool                   `json:"requireEmail,omitempty"`
	RequireJobTitle      bool                   `json:"requireJobTitle,omitempty"`
	RequireCompanyName   bool                   `json:"requireCompanyName,omitempty"`
	RequireAddress1      bool                   `json:"requireAddress1,omitempty"`
	RequireAddress2      bool                   `json:"requireAddress2,omitempty"`
	RequireCity          bool                   `json:"requireCity,omitempty"`
	RequireState         bool                   `json:"requireState,omitempty"`
	RequireZipCode       bool                   `json:"requireZipCode,omitempty"`
	RequireCountryRegion bool                   `json:"requireCountryRegion,omitempty"`
	RequireWorkPhone     bool                   `json:"requireWorkPhone,omitempty"`
	RequireFax           bool                   `json:"requireFax,omitempty"`
	MaxRegisterNum       int                    `json:"maxRegisterNum,omitempty"`
	CustomizedQuestions  []RegistrationQuestion `json:"customizedQuestions,omitempty"`
}

type RegistrationQuestion struct {
	QuestionID int      `json:"questionID,omitempty"`
	Question   string   `json:"question,omitempty"`
	Required   bool     `json:"required,omitempty"`
	Type       string   `json:"type,omitempty"`
	Options    []string `json:"options,omitempty"`
}

func (s *MeetingRegistrantsService) GetRegistrationForm(ctx context.Context, meetingID string, opts *QueryRegistrationFormOption) (*RegistrationForm, error) {
	if meetingID == "" {
		return nil, core.ErrInvalidParameter
	}

	params := url.Values{}
	if opts != nil {
		if opts.HostEmail != "" {
			params.Set("hostEmail", opts.HostEmail)
		}
		if opts.Current {
			params.Set("current", "true")
		}
	}

	var form RegistrationForm
	if err := s.session.Get(ctx, "meetings/"+meetingID+"/registration", params, &form); err != nil {
		return nil, err
	}
	return &form, nil
}
