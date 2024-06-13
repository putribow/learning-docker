package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/gin-gonic/gin"
)

func main() {
	// Get env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Connect to Database"
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("Connect to: ", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Connection to database failed: ", err)
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("ouch")
		return c.HTML(http.StatusOK,
			"if we never try, how would we know")
	})

	e.GET("/ping", func(c echo.Context) error {
		err := db.Ping()
		if err != nil {
			fmt.Println("PING Failed")
			return c.HTML(http.StatusOK,
				"if we never try, how would we know")
		}

		// Insert to database
		_, err = db.Exec("INSERT INTO timestamp (timestamp) VALUES ($1)", time.Now())
		if err != nil {
			fmt.Println("Failed to insert timestamp:", err)
			return c.HTML(http.StatusInternalServerError, "Ouch!")
		}

		fmt.Println("Ping successful")
		return c.HTML(http.StatusOK, "ping!")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "78"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
