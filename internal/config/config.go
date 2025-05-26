package config

import (
	"os"
	"strconv"
)

type Config struct {
	accessToken string
	baseURL     string
	timeOut     int
}

func NewConfig() *Config {
	accessToken := os.Getenv("WEBEX_ACCESS_TOKEN")
	baseUrl := os.Getenv("WEBEX_BASE_URL")
	timeOutStr := os.Getenv("WEBEX_TIMEOUT")

	if accessToken == "" {
		panic("WEBEX_ACCESS_TOKEN environment variable not set")
	}

	if baseUrl == "" {
		baseUrl = "https://webexapis.com"
	}

	var timeOut int
	if timeOutStr != "" {
		timeOutInt, err := strconv.Atoi(timeOutStr)
		if err != nil {
			timeOut = 30
		}
		timeOut = timeOutInt
	} else {
		timeOut = 30
	}

	return &Config{
		accessToken: accessToken,
		baseURL:     baseUrl,
		timeOut:     timeOut,
	}
}

func (c *Config) AccessToken() string {
	return c.accessToken
}

func (c *Config) BaseURL() string {
	return c.baseURL
}

func (c *Config) TimeOut() int {
	return c.timeOut
}
