package calling

import (
	"context"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type Call struct {
	ID             string       `json:"id,omitempty"`
	CallID         string       `json:"callId,omitempty"`
	CallSessionID  string       `json:"callSessionId,omitempty"`
	Personality    string       `json:"personality,omitempty"`
	State          string       `json:"state,omitempty"`
	RemoteParty    *RemoteParty `json:"remoteParty,omitempty"`
	Appearance     int          `json:"appearance,omitempty"`
	Created        time.Time    `json:"created,omitempty"`
	Connected      time.Time    `json:"connected,omitempty"`
	Duration       int          `json:"duration,omitempty"`
	Held           bool         `json:"held,omitempty"`
	RedirectReason string       `json:"redirectReason,omitempty"`
	RecordingState string       `json:"recordingState,omitempty"`
}

type RemoteParty struct {
	Name           string `json:"name,omitempty"`
	Number         string `json:"number,omitempty"`
	PersonID       string `json:"personId,omitempty"`
	PlaceID        string `json:"placeId,omitempty"`
	PrivacyEnabled bool   `json:"privacyEnabled,omitempty"`
	CallType       string `json:"callType,omitempty"`
}

type CallsService struct {
	session *core.RestSession
}

func NewCallsService(session *core.RestSession) *CallsService {
	return &CallsService{
		session: session,
	}
}

type DialRequest struct {
	Destination string `json:"destination"`
}

type DialResponse struct {
	CallID        string `json:"callId,omitempty"`
	CallSessionID string `json:"callSessionId,omitempty"`
}

func (s *CallsService) Dial(ctx context.Context, req *DialRequest) (*DialResponse, error) {
	if req == nil || req.Destination == "" {
		return nil, core.ErrInvalidParameter
	}

	var response DialResponse
	if err := s.session.Post(ctx, "telephony/calls/dial", req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type CallIDRequest struct {
	CallID string `json:"callId"`
}

func (s *CallsService) Answer(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/answer", req, nil)
}

func (s *CallsService) Reject(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/reject", req, nil)
}

func (s *CallsService) Hold(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/hold", req, nil)
}

func (s *CallsService) Resume(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/resume", req, nil)
}

func (s *CallsService) Hangup(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/hangup", req, nil)
}

type TransferRequest struct {
	CallID1     string `json:"callId1,omitempty"`
	CallID2     string `json:"callId2,omitempty"`
	CallID      string `json:"callId,omitempty"`
	Destination string `json:"destination,omitempty"`
}

func (s *CallsService) Transfer(ctx context.Context, callID string, destination string) error {
	if callID == "" || destination == "" {
		return core.ErrInvalidParameter
	}

	req := &TransferRequest{CallID: callID, Destination: destination}
	return s.session.Post(ctx, "telephony/calls/transfer", req, nil)
}

func (s *CallsService) ConsultTransfer(ctx context.Context, callID1, callID2 string) error {
	if callID1 == "" || callID2 == "" {
		return core.ErrInvalidParameter
	}

	req := &TransferRequest{CallID1: callID1, CallID2: callID2}
	return s.session.Post(ctx, "telephony/calls/consultTransfer", req, nil)
}

type DivertRequest struct {
	CallID      string `json:"callId"`
	Destination string `json:"destination,omitempty"`
	ToVoicemail bool   `json:"toVoicemail,omitempty"`
}

func (s *CallsService) Divert(ctx context.Context, callID string, destination string, toVoicemail bool) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	if !toVoicemail && destination == "" {
		return core.ErrInvalidParameter
	}

	req := &DivertRequest{CallID: callID, Destination: destination, ToVoicemail: toVoicemail}
	return s.session.Post(ctx, "telephony/calls/divert", req, nil)
}

type ParkRequest struct {
	CallID      string `json:"callId"`
	Destination string `json:"destination,omitempty"`
	IsGroupPark bool   `json:"isGroupPark,omitempty"`
}

type ParkResponse struct {
	ParkedAgainst string `json:"parkedAgainst,omitempty"`
}

func (s *CallsService) Park(ctx context.Context, callID string, destination string) (*ParkResponse, error) {
	if callID == "" {
		return nil, core.ErrInvalidParameter
	}

	req := &ParkRequest{
		CallID:      callID,
		Destination: destination,
	}

	var response ParkResponse
	if err := s.session.Post(ctx, "telephony/calls/park", req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type RetrieveRequest struct {
	Destination string `json:"destination"`
}

type RetrieveResponse struct {
	CallID        string `json:"callId,omitempty"`
	CallSessionID string `json:"callSessionId,omitempty"`
}

func (s *CallsService) Retrieve(ctx context.Context, destination string) (*RetrieveResponse, error) {
	if destination == "" {
		return nil, core.ErrInvalidParameter
	}

	req := &RetrieveRequest{Destination: destination}
	var response RetrieveResponse
	if err := s.session.Post(ctx, "telephony/calls/retrieve", req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type PickupRequest struct {
	Target string `json:"target"`
}

type PickupResponse struct {
	CallID        string `json:"callId,omitempty"`
	CallSessionID string `json:"callSessionId,omitempty"`
}

func (s *CallsService) Pickup(ctx context.Context, target string) (*PickupResponse, error) {
	if target == "" {
		return nil, core.ErrInvalidParameter
	}

	req := &PickupRequest{Target: target}
	var response PickupResponse
	if err := s.session.Post(ctx, "telephony/calls/pickup", req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type BargeInRequest struct {
	Target string `json:"target"`
}

type BargeInResponse struct {
	CallID        string `json:"callId,omitempty"`
	CallSessionID string `json:"callSessionId,omitempty"`
}

func (s *CallsService) BargeIn(ctx context.Context, target string) (*BargeInResponse, error) {
	if target == "" {
		return nil, core.ErrInvalidParameter
	}

	req := &BargeInRequest{Target: target}
	var response BargeInResponse
	if err := s.session.Post(ctx, "telephony/calls/bargein", req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *CallsService) StartRecording(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/startRecording", req, nil)
}

func (s *CallsService) StopRecording(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/stopRecording", req, nil)
}

func (s *CallsService) PauseRecording(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/pauseRecording", req, nil)
}

func (s *CallsService) ResumeRecording(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/resumeRecording", req, nil)
}

type TransmitDTMFRequest struct {
	CallID string `json:"callId"`
	DTMF   string `json:"dtmf"`
}

func (s *CallsService) TransmitDTMF(ctx context.Context, callID, dtmf string) error {
	if callID == "" || dtmf == "" {
		return core.ErrInvalidParameter
	}

	req := &TransmitDTMFRequest{CallID: callID, DTMF: dtmf}
	return s.session.Post(ctx, "telephony/calls/transmitDTMF", req, nil)
}

func (s *CallsService) Push(ctx context.Context, callID string) error {
	if callID == "" {
		return core.ErrInvalidParameter
	}

	req := &CallIDRequest{CallID: callID}
	return s.session.Post(ctx, "telephony/calls/push", req, nil)
}

func (s *CallsService) ListCalls(ctx context.Context) ([]*Call, error) {
	var response struct {
		Items []*Call `json:"items"`
	}

	if err := s.session.Get(ctx, "telephony/calls", nil, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (s *CallsService) GetCallDetails(ctx context.Context, callID string) (*Call, error) {
	if callID == "" {
		return nil, core.ErrInvalidParameter
	}

	var response Call
	if err := s.session.Get(ctx, "telephony/calls/"+callID, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
