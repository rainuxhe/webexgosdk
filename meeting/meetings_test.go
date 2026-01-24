package meeting

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

func TestMeetingsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/meetings" {
			t.Errorf("expected /meetings path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"items": [
				{
					"id": "meeting-123",
					"title": "Test Meeting",
					"meetingNumber": "1234567890",
					"meetingType": "meetingSeries"
				}
			]
		}`))
	}))
	defer server.Close()

	session := core.NewRestSession(&core.RestSessionConfig{
		AccessToken: "test-token",
		BaseURL:     server.URL + "/",
	})

	service := NewMeetingsService(session)
	meetings, err := service.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(meetings) != 1 {
		t.Errorf("expected 1 meeting, got %d", len(meetings))
	}
	if meetings[0].ID != "meeting-123" {
		t.Errorf("expected ID meeting-123, got %s", meetings[0].ID)
	}
	if meetings[0].Title != "Test Meeting" {
		t.Errorf("expected Title Test Meeting, got %s", meetings[0].Title)
	}
	if meetings[0].MeetingNumber != "1234567890" {
		t.Errorf("expected MeetingNumber 1234567890, got %s", meetings[0].MeetingNumber)
	}
	if meetings[0].MeetingType != "meetingSeries" {
		t.Errorf("expected MeetingType meetingSeries, got %s", meetings[0].MeetingType)
	}
}

func TestMeetingsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/meetings" {
			t.Errorf("expected /meetings path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{
			"id": "meeting-456",
			"title": "New Meeting",
			"meetingNumber": "9876543210",
			"meetingType": "meetingSeries",
			"webLink": "https://webex.com/join/meeting-456"
		}`))
	}))
	defer server.Close()

	session := core.NewRestSession(&core.RestSessionConfig{
		AccessToken: "test-token",
		BaseURL:     server.URL + "/",
	})

	service := NewMeetingsService(session)
	now := time.Now()
	meeting, err := service.Create(context.Background(), &MeetingCreateRequest{
		MeetingRequestBase: MeetingRequestBase{
			Title: "New Meeting",
			Start: now,
			End:   now.Add(time.Hour),
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if meeting.ID != "meeting-456" {
		t.Errorf("expected ID meeting-456, got %s", meeting.ID)
	}
	if meeting.Title != "New Meeting" {
		t.Errorf("expected Title New Meeting, got %s", meeting.Title)
	}
	if meeting.MeetingNumber != "9876543210" {
		t.Errorf("expected MeetingNumber 9876543210, got %s", meeting.MeetingNumber)
	}
	if meeting.MeetingType != "meetingSeries" {
		t.Errorf("expected MeetingType meetingSeries, got %s", meeting.MeetingType)
	}
}
