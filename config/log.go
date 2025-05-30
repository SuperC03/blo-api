package config

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(
	isProd bool,
	gotify *GotifyClient,
) (*zap.Logger, error) {
	var cfg zap.Config
	if isProd {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.WithOptions(zap.Hooks(func(e zapcore.Entry) error {
		if e.Level >= zap.ErrorLevel {
			// In theory should a goroutine, but if we're here, something's already messed up :(
			gotify.Send(context.Background(), e.Level.String(), e.Message)
		}
		return nil
	})), nil
}
