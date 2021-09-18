package route

import (
	"fmt"
	"net/http"
	"os"

	"piovani/golang_with_kafka/controller"

	"github.com/labstack/echo/v4"
)

func probe(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("service is running on port %s", os.Getenv("APP_PORT")))
}

func InitRoute() {
	e := echo.New()

	e.GET("/probe", probe)

	e.GET("/read", controller.Read())
	// e.GET("/publish")

	e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
