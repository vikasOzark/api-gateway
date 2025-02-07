package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

/*
This function is responsible to handle the response's according to the
response's content-type this function will return the response to the client.
*/
func ContentTypeResponse(c echo.Context, contentType string, data []byte) error {
	contentType = strings.Split(contentType, ";")[0]

	switch contentType {
	case "application/json":
		var jsonResponse map[string]interface{}
		if err := json.Unmarshal(data, &jsonResponse); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid JSON response"})
		}
		return c.JSON(http.StatusOK, jsonResponse)

	case "text/html", "text/plain":
		return c.HTML(http.StatusOK, string(data))

	case contentType:
		return c.XML(http.StatusOK, data)

	case "application/octet-stream", "application/pdf", "image/png", "image/jpeg", "image/gif":
		return c.Blob(http.StatusOK, contentType, data)

	default:
		return c.String(http.StatusUnsupportedMediaType, "Unsupported content type")
	}
}
