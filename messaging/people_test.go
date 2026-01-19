package messaging

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

func TestPeopleServiceGetMe(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		
		if r.URL.Path != "/people/me" {
			t.Errorf("expected /people/me path, got %s", r.URL.Path)
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "person-123",
			"emails": ["test1@example.com"],
			"displayName": "Test User",
			"firstName": "Test",
			"lastName": "User"
			}`))
	}))
	defer server.Close()
	
	session := core.NewRestSession(&core.RestSessionConfig{
		AccessToken: "test-token",
		BaseURL: server.URL+"/",
	})
	
	service := NewPeopleService(session)
	person, err := service.GetMe(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	
	if person.ID != "person-123" {
		t.Errorf("expected ID person-123, got %s", person.ID)
	}
	
	if person.Emails[0] != "test1@example.com" {
		t.Errorf("expected email test1@example.com, got %s", person.Emails[0])
	}
	
	if person.DisplayName != "Test User" {
		t.Errorf("expected display name Test User, got %s", person.DisplayName)
	}
	
	if person.FirstName != "Test" {
		t.Errorf("expected first name Test, got %s", person.FirstName)
	}
	
	if person.LastName != "User" {
		t.Errorf("expected last name User, got %s", person.LastName)
	}
}