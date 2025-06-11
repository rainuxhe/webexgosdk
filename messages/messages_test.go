package messages

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	messages := NewMessages()
	resp, err := messages.Create(
		WithRoomId(os.Getenv("TEST_ROOM_ID")),
		WithMarkdown("Hello from webexgosdk"),
	)
	if err != nil {
		t.Errorf("Error creating message: %v", err)
	} else {
		t.Logf("Message created successfully, %+v\n", resp)
	}

}

func TestGet(t *testing.T) {
	messageId := os.Getenv("TEST_MESSAGE_ID")
	messages := NewMessages()
	resp, err := messages.Get(messageId)
	if err != nil {
		t.Errorf("Error getting message: %v", err)
	} else {
		t.Logf("Message get successfully, %+v\n", resp)
	}
}

func TestListDirect(t *testing.T) {
	personEmail := os.Getenv("TEST_PERSON_EMAIL")
	messages := NewMessages()
	resp, err := messages.ListDirect(WithToPersonEmail(personEmail))
	if err != nil {
		t.Errorf("Error listing direct messages: %v", err)
	} else {
		t.Logf("Direct messages list successfully, %+v\n", resp)
	}
}
