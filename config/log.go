package config

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(
	cfg EnvConfig,
	gotify *GotifyClient,
) (*zap.Logger, error) {
	var logCfg zap.Config
	if cfg.Production {
		logCfg = zap.NewProductionConfig()
	} else {
		logCfg = zap.NewDevelopmentConfig()
	}

	logCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logCfg.EncoderConfig.TimeKey = "timestamp"
	logCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := logCfg.Build()
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
