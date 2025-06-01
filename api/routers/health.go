package routers

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/superc03/blo-api/config"
	"go.uber.org/zap"
)

type healthRouter = AppRouter

func NewHealthRouter(
	g *echo.Group,
	db *pgxpool.Pool,
	log *zap.Logger,
	cfg config.EnvConfig,
	gotify *config.GotifyClient,
	sheets *config.SheetsClient,
) *AppRouter {
	r := &healthRouter{
		router: *g,
		db:     db,
		log:    log,
		cfg:    cfg,
		gotify: gotify,
		sheets: sheets,
	}

	r.router.GET("/", r.getHealth)
	return r
}

// getHealth godoc
// @Summary Indicates server health
// @Tags meta
// @Produce json
// @Success 200 {object} routers.GenericJsonDto
// @Failure 500 {object} routers.GenericJsonDto
// @Router /health [get]
func (r *healthRouter) getHealth(c echo.Context) error {
	err := r.db.Ping(c.Request().Context())
	if err == nil {
		return c.JSON(http.StatusOK, GenericJsonDto{
			Status: "success",
			Data:   "okie dokie",
		})
	} else {
		r.log.Error("database error from health check", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, GenericJsonDto{
			Status: "error",
			Data:   "not okie dokie",
		})
	}
}
