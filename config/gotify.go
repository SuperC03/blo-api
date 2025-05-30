package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type GotifyClient struct {
	enabled bool
	url     string
}

func (
	g *GotifyClient) Send(
	ctx context.Context,
	title string,
	body string,
) error {
	if !g.enabled {
		return nil
	}

	data := url.Values{}
	data.Add("title", title)
	data.Add("message", body)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, g.url, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("error reaching Gotify server")
	}

	return nil
}

func NewGotifyClient(
	enabled bool,
	host string,
	appToken string,
) *GotifyClient {
	return &GotifyClient{
		enabled: enabled,
		url:     fmt.Sprintf("%s/message?token=%s", host, appToken),
	}
}
