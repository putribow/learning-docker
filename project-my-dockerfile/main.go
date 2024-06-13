package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/gin-gonic/gin"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("ouch")
		return c.HTML(http.StatusOK,
			"if we never try, how would we know")
	})
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "77"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
