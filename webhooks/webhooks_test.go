package webhooks

import "testing"

func TestList(t *testing.T) {
	webhooks := NewWebhooks()
	webhooksList, err := webhooks.List()
	if err != nil {
		t.Errorf("Error listing webhooks: %v", err)
	} else {
		t.Logf("Webhooks: %+v", webhooksList)
	}
}
