package config

import (
	"context"
	"os"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetsClient struct {
	config jwt.Config
}

func (s *SheetsClient) GetRange(
	ctx context.Context,
	spreadsheetId string,
	range_ string,
) (*sheets.ValueRange, error) {
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(s.config.Client(ctx)))
	if err != nil {
		return nil, err
	}

	return srv.Spreadsheets.Values.Get(spreadsheetId, range_).Do()
}

func NewSheetsClient(
	cfg EnvConfig,
) (*SheetsClient, error) {
	creds, err := os.ReadFile(cfg.GoogleKeyPath)
	if err != nil {
		return nil, err
	}

	scopes := []string{
		"https://www.googleapis.com/auth/spreadsheets.readonly",
	}

	config, err := google.JWTConfigFromJSON(creds, scopes...)
	if err != nil {
		return nil, err
	}

	return &SheetsClient{
		config: *config,
	}, nil
}
