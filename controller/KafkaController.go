package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Read(c echo.Context) error {
	return echo.HandlerFunc(c.JSON(http.StatusOK, "READ"))
}
