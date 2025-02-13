package test

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func EnvSetter(t *testing.T) {
	t.Setenv("ISDELVE_Enabled", "true")
}

func EchoRouterProvider(t *testing.T) (*echo.Echo, *assert.Assertions) {
	EnvSetter(t)
	assert := assert.New(t)
	e := echo.New()
	return e, assert
}
