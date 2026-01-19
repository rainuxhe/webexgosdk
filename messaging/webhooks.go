package messaging

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
)

type Webhook struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	TargetURL string    `json:"targetUrl,omitempty"`
	Resource  string    `json:"resource,omitempty"`
	Event     string    `json:"event,omitempty"`
	OrgID     string    `json:"orgId,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty"`
	AppID     string    `json:"appId,omitempty"`
	OwnedBy   string    `json:"ownedBy,omitempty"`
	Filter    string    `json:"filter,omitempty"`
	Secret    string    `json:"secret,omitempty"`
	Status    string    `json:"status,omitempty"`
	Created   time.Time `json:"created,omitempty"`
}

type WebhooksService struct {
	session *core.RestSession
}

func NewWebhooksService(session *core.RestSession) *WebhooksService {
	return &WebhooksService{
		session: session,
	}
}

type WebhookListOptions struct {
	Max int
}

func (s *WebhooksService) List(ctx context.Context, opts *WebhookListOptions) ([]*Webhook, error) {
	params := url.Values{}
	if opts != nil {
		if opts.Max > 0 {
			params.Set("max", strconv.Itoa(opts.Max))
		}
	}

	var response struct {
		Items []*Webhook `json:"items"`
	}

	if err := s.session.Get(ctx, "webhooks", params, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}

type WebhookCreateRequest struct {
	Name      string `json:"name"`
	TargetURL string `json:"targetUrl"`
	Resource  string `json:"resource"`
	Event     string `json:"event"`
	Filter    string `json:"filter,omitempty"`
	Secret    string `json:"secret,omitempty"`
	OwnedBy   string `json:"ownedBy,omitempty"`
}

func (s *WebhooksService) Create(ctx context.Context, req *WebhookCreateRequest) (*Webhook, error) {
	if req == nil || req.Name == "" || req.TargetURL == "" || req.Resource == "" || req.Event == "" {
		return nil, core.ErrInvalidParameter
	}

	var webhook Webhook
	if err := s.session.Post(ctx, "webhooks", req, &webhook); err != nil {
		return nil, err
	}

	return &webhook, nil
}

func (s *WebhooksService) Get(ctx context.Context, webhookID string) (*Webhook, error) {
	if webhookID == "" {
		return nil, core.ErrInvalidParameter
	}

	var webhook Webhook
	if err := s.session.Get(ctx, "webhook/"+webhookID, nil, &webhook); err != nil {
		return nil, err
	}

	return &webhook, nil
}

type WebhookUpdateRequest struct {
	Name      string `json:"name"`
	TargetURL string `json:"targetUrl"`
	Secret    string `json:"secret,omitempty"`
	OwnedBy   string `json:"ownedBy,omitempty"`
	Status    string `json:"status,omitempty"`
}

func (s *WebhooksService) Update(ctx context.Context, webhookID string, req *WebhookUpdateRequest) (*Webhook, error) {
	if webhookID == "" || req == nil || req.Name == "" || req.TargetURL == "" {
		return nil, core.ErrInvalidParameter
	}

	var webhook Webhook
	if err := s.session.Put(ctx, "webhooks/"+webhookID, req, &webhook); err != nil {
		return nil, err
	}

	return &webhook, nil
}

func (s *WebhooksService) Delete(ctx context.Context, webhookID string) error {
	if webhookID == "" {
		return core.ErrInvalidParameter
	}

	return s.session.Delete(ctx, "webhooks/"+webhookID)
}
