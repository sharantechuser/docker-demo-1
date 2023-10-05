package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sharantechuser/docker1/config"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":4000"))
}

// Handler
func hello(c echo.Context) error {
	log.Println("hello...")
	out_msg := fmt.Sprintf("Hello %s!", config.User)
	return c.String(http.StatusOK, out_msg)
}
