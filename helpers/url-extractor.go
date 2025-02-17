package helpers

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// UrlExtractor is a struct that extracts and processes URL components from an Echo context.
type UrlExtractor struct {
    URL     string       // The full URL path
    Service string       // The service name extracted from the URL
    HOST    string       // The host of the request
    URI     string       // The URI part of the URL
    QUERY   string       // The query string of the URL
    RawURL  string       // The raw URL string
    Method  string       // The HTTP method of the request
    Context echo.Context // The Echo context
}

// ExtractURL splits the URL path into its components and returns them as a slice of strings.
func (u *UrlExtractor) ExtractURL() []string {
    path := u.Context.Request().URL.Path
    urlPath := strings.Split(path, "/")
    return urlPath
}

// Process extracts and processes the URL components from the Echo context and populates the UrlExtractor fields.
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

// GetTargetUrl constructs the target URL for the service by loading the configuration and appending the URI.
func (u *UrlExtractor) GetTargetUrl() string {
    config := &Config{}
    config.LoadConfig()
    return config.GetTarget(u.Service) + u.URI
}
