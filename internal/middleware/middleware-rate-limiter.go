package middleware

import (
	"gateway/helpers"
	"gateway/helpers/constant"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func RateLimiterConfig() middleware.RateLimiterConfig {
	RATE_LIMIT_BURST_REQUEST, _ := helpers.ConvertEnvInt(constant.ENV_RATE_LIMIT_BURST_REQUEST)
	RATE_LIMIT_REQ_PER_SEC, _ := helpers.ConvertEnvInt(constant.ENV_RATE_LIMIT_REQ_PER_SEC)

	Config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(RATE_LIMIT_REQ_PER_SEC), Burst: RATE_LIMIT_BURST_REQUEST, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	return Config
}
