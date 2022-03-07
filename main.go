package main

import (
	"fmt"
	"os"

	_ "net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	port = getEnvOrDefault("HTTPS_PORT", "8443")
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("assets"))
	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

// getEnvOrDefault returns the value for the env var or a default if the env var is not set.
func getEnvOrDefault(key, def string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return def
}
