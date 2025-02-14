package test

import (
	"gateway/helpers"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLoader() {
	helpers.RequiredChecks()
	config := &helpers.Config{}
	config.LoadConfig()
}

func EnvSetter(t *testing.T) {
	TestLoader()
}

func EchoRouterProvider(t *testing.T) (*echo.Echo, *assert.Assertions) {
	EnvSetter(t)
	assert := assert.New(t)
	e := echo.New()
	return e, assert
}
