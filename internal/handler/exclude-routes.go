package handler

import (
	"gateway/helpers"
	"gateway/internal/requests"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
)

func ExcludeRouteHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := helpers.UrlExtractor{
			Context: c,
		}
		url.Process()

		config := &helpers.Config{}
		config.LoadConfig()

		value, ok := config.IsExcluded(url.Service)
		if !ok {
			return next(c)
		}

		excludedUrls := strings.Split(value, ",")

		urls := []string{}
		for _, v := range excludedUrls {
			str := []string{strings.TrimSpace(v)}
			urls = append(urls, str...)
		}

		isThere := slices.Contains(urls, url.URL)

		if !isThere {
			return next(c)
		}
		return requests.RequestManager(url)
	}
}
