package webhooks

import "time"

type Response struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	TargetURL string    `json:"targetUrl"`
	Resource  string    `json:"resource"`
	Event     string    `json:"event"`
	Filter    string    `json:"filter"`
	Secret    string    `json:"secret"`
	Status    string    `json:"status"`
	Created   time.Time `json:"created"`
	OwnedBy   string    `json:"ownedBy"`
}
