package webexgosdk

import (
	"errors"
	"os"
	"time"

	"github.com/rainuxhe/webexgosdk/internal/core"
	"github.com/rainuxhe/webexgosdk/messaging"
)

const (
	DefaultBaseURL        = "https://webexapis.com/v1/"
	DefaultRequestTimeout = 60
	AccessTokenEnvVar     = "WEBEX_ACCESS_TOKEN"
	Version               = "0.0.1"
)

var (
	ErrAccessTokenRequired = errors.New("access token is required")
)

type MessagingAPI struct {
	Messages *messaging.MessagesService
	// Rooms *messaging.RoomsService
}

type Client struct {
	Messaging *MessagingAPI

	session *core.RestSession
}

type ClientOptions struct {
	BaseURL        string
	RequestTimeout time.Duration
	UserAgent      string
}

func NewClient(accessToken string, opts ...ClientOptions) (*Client, error) {
	if accessToken == "" {
		accessToken = os.Getenv(AccessTokenEnvVar)
	}

	if accessToken == "" {
		return nil, ErrAccessTokenRequired
	}

	options := ClientOptions{
		BaseURL:        DefaultBaseURL,
		RequestTimeout: DefaultRequestTimeout * time.Second,
	}

	if len(opts) > 0 {
		opt := opts[0]
		if opt.BaseURL != "" {
			options.BaseURL = opt.BaseURL
		}

		if opt.RequestTimeout > 0 {
			options.RequestTimeout = opt.RequestTimeout
		}

		if opt.UserAgent != "" {
			options.UserAgent = opt.UserAgent
		}
	}

	session := core.NewRestSession(&core.RestSessionConfig{
		AccessToken: accessToken,
		BaseURL:     options.BaseURL,
		Timeout:     options.RequestTimeout,
		UserAgent:   options.UserAgent,
	})

	client := &Client{
		session: session,
	}

	client.Messaging = &MessagingAPI{
		Messages: messaging.NewMessagesService(session),

		// ...
	}

	return client, nil
}

func (c *Client) AccessToken() string {
	return c.session.GetAccessToken()
}

func (c *Client) SetAccessToken(token string) {
	c.session.SetAccessToken(token)
}
