package routers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/superc03/blo-api/config"
	"go.uber.org/zap"
)

type AppRouter struct {
	router echo.Group
	db     *pgxpool.Pool
	log    *zap.Logger
	cfg    config.EnvConfig
	gotify *config.GotifyClient
	sheets *config.SheetsClient
}

// Roughly inspired from https://github.com/omniti-labs/jsend
type GenericJsonDto struct {
	Status  string `json:"status" example:"success"`
	Data    string `json:"data" example:"okie dokie"`
	Message string `json:"message" example:"internal server error"`
}
