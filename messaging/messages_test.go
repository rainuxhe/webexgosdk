package messaging

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

func TestMessagesServiceCreate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("excepted POST, got %s", r.Method)
		}
		if r.URL.Path != "/messages" {
			t.Errorf("expected /messages, got %s", r.URL.Path)
		}

		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-token" {
			t.Errorf("expected Bearer test-token, got %s", auth)
		}
		resp := Message{
			ID:     "test-message-id",
			RoomID: "test-room-id",
			Text:   "Hello World",
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	session := core.NewRestSession(&core.RestSessionConfig{
		AccessToken: "test-token",
		BaseURL: server.URL + "/",
	})

	ctx := context.Background()
	service := NewMessagesService(session)
	message, err := service.Create(ctx, &MessageCreateRequest{
		RoomID: "test-room-id",
		Text: "Hello World",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if message.ID != "test-message-id" {
		t.Errorf("expected test-message-id, got %s", err)
	}
}
