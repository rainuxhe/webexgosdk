package meeting

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type Meeting struct {
	ID            string    `json:"id,omitempty"`
	MeetingNumber string    `json:"meetingNumber,omitempty"`
	Title         string    `json:"title,omitempty"`
	Agenda        string    `json:"agenda,omitempty"`
	Password      string    `json:"password,omitempty"`
	MeetingType   string    `json:"meetingType,omitempty"`
	State         string    `json:"state,omitempty"`
	Timezone      string    `json:"timezone,omitempty"`
	Start         time.Time `json:"start,omitempty"`
	End           time.Time `json:"end,omitempty"`

	// Recurrence is the recurrence pattern in iCalendar PRULE format.
	Recurrence string `json:"recurrence,omitempty"`

	// HostUserID is the unique identifier of the host user.
	HostUserID string `json:"hostUserId,omitempty"`

	HostDisplayName           string `json:"hostDisplayName,omitempty"`
	HostEmail                 string `json:"hostEmail,omitempty"`
	HostKey                   string `json:"hostKey,omitempty"`
	SiteURL                   string `json:"siteUrl,omitempty"`
	WebLink                   string `json:"webLink,omitempty"`
	SipAddress                string `json:"sipAddress,omitempty"`
	DialInIPAddress           string `json:"dialInIpAddress,omitempty"`
	EnabledAutoRecordMeeting  bool   `json:"enabledAutoRecordMeeting,omitempty"`
	AllowAnyUserToBeCoHost    bool   `json:"allowAnyUserToBeCoHost,omitempty"`
	AllowFirstUserToBeCoHost  bool   `json:"allowFirstUserToBeCoHost,omitempty"`
	AllowAuthenticatedDevices bool   `json:"allowAuthenticatedDevices,omitempty"`
	EnabledJoinBeforeHost     bool   `json:"enabledJoinBeforeHost,omitempty"`

	// JoinBeforeHostMinutes is the minutes before host can join.
	JoinBeforeHostMinutes int `json:"joinBeforeHostMinutes,omitempty"`

	EnableConnectAudioBeforeHost        bool                        `json:"enableConnectAudioBeforeHost,omitempty"`
	ExcludePassword                     bool                        `json:"excludePassword,omitempty"`
	PublicMeeting                       bool                        `json:"publicMeeting,omitempty"`
	ReminderTime                        int                         `json:"reminderTime,omitempty"`
	UnlockedMeetingJoinSecurity         string                      `json:"unlockedMeetingJoinSecurity,omitempty"`
	SessionTypeID                       int                         `json:"sessionTypeId,omitempty"`
	ScheduledType                       string                      `json:"scheduledType,omitempty"`
	EnabledWebcastView                  bool                        `json:"enabledWebcastView,omitempty"`
	PanelistPassword                    string                      `json:"panelistPassword,omitempty"`
	PhoneAndVideoSystemPanelistPassword string                      `json:"phoneAndVideoSystemPanelistPassword,omitempty"`
	EnableAutomaticLockMinutes          int                         `json:"enableAutomaticLockMinutes,omitempty"`
	MeetingSeriesID                     string                      `json:"meetingSeriesId,omitempty"`
	ScheduleMeetingID                   string                      `json:"scheduleMeetingId,omitempty"`
	MeetingOptions                      *MeetingOptions             `json:"meetingOptions,omitempty"`
	AudioConnectionOptions              *AudioConnectionOptions     `json:"audioConnectionOptions,omitempty"`
	IntegrationTags                     []string                    `json:"integrationTags,omitempty"`
	Invitees                            []Invitee                   `json:"invitees,omitempty"`
	Registration                        *Registration               `json:"registration,omitempty"`
	SimultaneousInterpretation          *SimultaneousInterpretation `json:"simultaneousInterpretation,omitempty"`
	Breakout                            *BreakoutSessions           `json:"breakout,omitempty"`
}

type MeetingOptions struct {
	EnabledChat          bool   `json:"enabledChat,omitempty"`
	EnabledVideo         bool   `json:"enabledVideo,omitempty"`
	EnabledPolling       bool   `json:"enabledPolling,omitempty"`
	EnabledNote          bool   `json:"enabledNote,omitempty"`
	NoteType             string `json:"noteType,omitempty"`
	EnabledClosedCaption bool   `json:"enabledClosedCaption,omitempty"`
	EnabledFileTransfer  bool   `json:"enabledFileTransfer,omitempty"`
	EnabledUCFRichMedia  bool   `json:"enabledUCFRichMedia,omitempty"`
}

type AudioConnectionOptions struct {
	AudioConnectionType           string `json:"audioConnectionType,omitempty"`
	EnabledTollFreeCallIn         bool   `json:"enabledTollFreeCallIn,omitempty"`
	EnabledGlobalCallIn           bool   `json:"enabledGlobalCallIn,omitempty"`
	EnabledAudienceCallBack       bool   `json:"enabledAudienceCallBack,omitempty"`
	EntryAndExitTone              string `json:"entryAndExitTone,omitempty"`
	AllowHostToUnmuteParticipants bool   `json:"allowHostToUnmuteParticipants,omitempty"`
	AllowAttendeeToUnmuteSelf     bool   `json:"allowAttendeeToUnmuteSelf,omitempty"`
	MuteAttendeeUponEntry         bool   `json:"muteAttendeeUponEntry,omitempty"`
}

type Registration struct {
	AutoRegisterEnabled  bool   `json:"autoRegisterEnabled,omitempty"`
	AutoAcceptRequest    bool   `json:"autoAcceptRequest,omitempty"`
	RequireFirstName     bool   `json:"requireFirstName,omitempty"`
	RequireLastName      bool   `json:"requireLastName,omitempty"`
	RequireEmail         bool   `json:"requireEmail,omitempty"`
	RequireJobTitle      bool   `json:"requireJobTitle,omitempty"`
	RequireCompanyName   bool   `json:"requireCompanyName,omitempty"`
	RequireAddress1      bool   `json:"requireAddress1,omitempty"`
	RequireAddress2      bool   `json:"requireAddress2,omitempty"`
	RequireCity          bool   `json:"requireCity,omitempty"`
	RequireState         bool   `json:"requireState,omitempty"`
	RequireZipCode       bool   `json:"requireZipCode,omitempty"`
	RequireCountryRegion bool   `json:"requireCountryRegion,omitempty"`
	RequireWorkPhone     bool   `json:"requireWorkPhone,omitempty"`
	RequireFax           bool   `json:"requireFax,omitempty"`
	MaxRegisterNum       int    `json:"maxRegisterNum,omitempty"`
	RegisterFormID       string `json:"registerFormID,omitempty"`
}

type SimultaneousInterpretation struct {
	Enabled      bool          `json:"enabled,omitempty"`
	Interpreters []Interpreter `json:"interpreters,omitempty"`
}

type Interpreter struct {
	Email         string `json:"email,omitempty"`
	DisplayName   string `json:"displayName,omitempty"`
	LanguageCode1 string `json:"languageCode1,omitempty"`
	LanguageCode2 string `json:"languageCode2,omitempty"`
}

type BreakoutSessions struct {
	PreAssignmentes *PreAssignments `json:"preAssignments,omitempty"`
}

type PreAssignments struct {
	Enabled bool `json:"enabled,omitempty"`
}

type Invitee struct {
	Email       string `json:"email,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	CoHost      bool   `json:"coHost,omitempty"`
	Panelist    bool   `json:"panelist,omitempty"`
}

type MeetingsService struct {
	session *core.RestSession
}

func NewMeetingsService(session *core.RestSession) *MeetingsService {
	return &MeetingsService{
		session: session,
	}
}

type MeetingListOptions struct {
	MeetingNumber    string
	WebLink          string
	RoomID           string
	MeetingType      string
	State            string
	ScheduledType    string
	ParticipantEmail string
	Current          bool
	From             time.Time
	To               time.Time

	// HostEmail filters by host email(admin only)
	HostEmail      string
	SiteURl        string
	IntegrationTag string
	Max            int
}

func (s *MeetingsService) List(ctx context.Context, opts *MeetingListOptions) ([]*Meeting, error) {
	params := url.Values{}

	if opts != nil {
		if opts.MeetingNumber != "" {
			params.Set("meetingNumber", opts.MeetingNumber)
		}
		if opts.WebLink != "" {
			params.Set("webLink", opts.WebLink)
		}
		if opts.RoomID != "" {
			params.Set("roomId", opts.RoomID)
		}
		if opts.MeetingType != "" {
			params.Set("meetingType", opts.MeetingType)
		}
		if opts.State != "" {
			params.Set("state", opts.State)
		}
		if opts.ScheduledType != "" {
			params.Set("scheduledType", opts.ScheduledType)
		}
		if opts.ParticipantEmail != "" {
			params.Set("participantEmail", opts.ParticipantEmail)
		}
		if opts.Current {
			params.Set("current", "true")
		}
		if !opts.From.IsZero() {
			params.Set("from", opts.From.Format(time.RFC3339))
		}
		if !opts.To.IsZero() {
			params.Set("to", opts.To.Format(time.RFC3339))
		}
		if opts.HostEmail != "" {
			params.Set("hostEmail", opts.HostEmail)
		}
		if opts.SiteURl != "" {
			params.Set("siteUrl", opts.SiteURl)
		}
		if opts.IntegrationTag != "" {
			params.Set("integrationTag", opts.IntegrationTag)
		}
		if opts.Max > 0 {
			params.Set("max", strconv.Itoa(opts.Max))
		}
	}

	var response struct {
		Items []*Meeting `json:"items"`
	}

	if err := s.session.Get(ctx, "meetings", params, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

type MeetingRequestBase struct {
	Title                        string                  `json:"title"`
	Agenda                       string                  `json:"agenda,omitempty"`
	Password                     string                  `json:"password,omitempty"`
	Start                        time.Time               `json:"start"`
	End                          time.Time               `json:"end"`
	Timezone                     string                  `json:"timezone,omitempty"`
	Recurrence                   string                  `json:"recurrence,omitempty"`
	EnabledAutoRecordMeeting     bool                    `json:"enabledAutoRecordMeeting"`
	AllowAnyUserToBeCoHost       bool                    `json:"allowAnyUserToBeCoHost"`
	AllowFirstUserToBeCoHost     bool                    `json:"allowFirstUserToBeCoHost"`
	AllowAuthenticateDevices     bool                    `json:"allowAuthenticateDevices"`
	EnabledJoinBeforeHost        bool                    `json:"enabledJoinBeforeHost"`
	JoinBeforeHostMinutes        int                     `json:"joinBeforeHostMinutes"`
	EnableConnectAudioBeforeHost bool                    `json:"enableConnectAudioBeforeHost"`
	ExcludePassword              bool                    `json:"excludePassword,omitempty"`
	PublicMeeting                bool                    `json:"publicMeeting,omitempty"`
	ReminderTime                 int                     `json:"reminderTime,omitempty"`
	UnlockedMeetingJoinSecurity  string                  `json:"unlockedMeetingJoinSecurity,omitempty"`
	SessionTypeID                int                     `json:"sessionTypeId,omitempty"`
	ScheduledType                string                  `json:"scheduledType,omitempty"`
	EnabledWebcastView           bool                    `json:"enabledWebcastView"`
	PanelistPassword             string                  `json:"panelistPassword,omitempty"`
	EnableAutomaticLock          bool                    `json:"enableAutomaticLock"`
	AutomaticLockMinutes         int                     `json:"automaticLockMinutes"`
	HostEmail                    string                  `json:"hostEmail,omitempty"`
	SiteURL                      string                  `json:"siteUrl,omitempty"`
	MeetingOptions               *MeetingOptions         `json:"meetingOptions,omitempty"`
	AudioConnectionOptions       *AudioConnectionOptions `json:"audioConnectionOptions,omitempty"`
	IntegrationTags              []string                `json:"integrationTags,omitempty"`
	SendEmail                    bool                    `json:"sendEmail,omitempty"`
}

type MeetingCreateRequest struct {
	MeetingRequestBase
	Invitees     []Invitee     `json:"invitees,omitempty"`
	Registration *Registration `json:"registration,omitempty"`
}

func (s *MeetingsService) Create(ctx context.Context, req *MeetingCreateRequest) (*Meeting, error) {
	if req == nil || req.Title == "" || req.Start.IsZero() || req.End.IsZero() {
		return nil, core.ErrInvalidParameter
	}

	var meeting Meeting
	if err := s.session.Post(ctx, "meetings", req, &meeting); err != nil {
		return nil, err
	}
	return &meeting, nil
}

// Get returns details of a meeting by ID.
func (s *MeetingsService) Get(ctx context.Context, meetingID string) (*Meeting, error) {
	if meetingID == "" {
		return nil, core.ErrInvalidParameter
	}

	var meeting Meeting
	if err := s.session.Get(ctx, "meetings/"+meetingID, nil, &meeting); err != nil {
		return nil, err
	}

	return &meeting, nil
}

type MeetingUpdateRequest struct {
	MeetingRequestBase
}

func (s *MeetingsService) Update(ctx context.Context, meetingID string, req *MeetingUpdateRequest) (*Meeting, error) {
	if meetingID == "" || req == nil || req.Title == "" || req.Start.IsZero() || req.End.IsZero() {
		return nil, core.ErrInvalidParameter
	}

	var meeting Meeting
	if err := s.session.Put(ctx, "meetings/"+meetingID, req, &meeting); err != nil {
		return nil, err
	}
	return &meeting, nil
}

func (s *MeetingsService) Delete(ctx context.Context, meetingID string) error {
	if meetingID == "" {
		return core.ErrInvalidParameter
	}

	return s.session.Delete(ctx, "meetings/"+meetingID)
}

type MeetingJoinInfo struct {
	JoinLink       string    `json:"joinLink,omitempty"`
	ExpirationTime time.Time `json:"expirationTime,omitempty"`
}

func (s *MeetingsService) Join(ctx context.Context, meetingID string, email string, displayName string) (*MeetingJoinInfo, error) {
	if meetingID == "" {
		return nil, core.ErrInvalidParameter
	}

	reqBody := map[string]any{
		"meetingId": meetingID,
	}
	if email != "" {
		reqBody["email"] = email
	}
	if displayName != "" {
		reqBody["displayName"] = displayName
	}

	var joinInfo MeetingJoinInfo
	if err := s.session.Post(ctx, "meetings/join", reqBody, &joinInfo); err != nil {
		return nil, err
	}
	return &joinInfo, nil
}
