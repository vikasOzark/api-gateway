package helpers

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type UrlExtractor struct {
	URL     string
	Service string
	HOST    string
	URI     string
	QUERY   string
	RawURL  string
	Method  string
	Context echo.Context
}

func (u *UrlExtractor) ExtractURL() []string {
	path := u.Context.Request().URL.Path
	urlPath := strings.Split(path, "/")
	return urlPath
}

func (u *UrlExtractor) Process() {
	context := u.Context
	path := context.Request().URL.Path
	u.URL = path

	u.HOST = context.Request().Host
	u.QUERY = context.Request().URL.RawQuery
	u.RawURL = context.Request().URL.String()
	u.Method = context.Request().Method

	urlPath := strings.Split(u.RawURL, "/")
	if len(urlPath) > 0 && urlPath[0] == "" {
		urlPath = urlPath[1:]
	}

	var segments []string
	for _, segment := range urlPath {
		subSegments := strings.Fields(segment)
		segments = append(segments, subSegments...)
	}

	u.Service = segments[0]
	u.URI = "/" + strings.Join(segments[1:], "/")
}

func (u *UrlExtractor) GetTargetUrl() string {
	config := &Config{}
	config.LoadConfig()
	return config.GetTarget(u.Service) + u.URI
}
