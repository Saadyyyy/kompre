package main

import (
	// "bank_soal/config"
	// "bank_soal/route"
	"fmt"
	"kompre/config"
	route "kompre/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()

	// Initialize GORM DB
	gormDB, err := config.InitDBPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform migrations
	config.DBMigration(gormDB)

	// Initialize Echo instance
	e := echo.New()

	// Middleware setup
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Register routes with Echo
	route.Register(gormDB, e)

	port := fmt.Sprintf(":%d", cfg.SERVERPORT)
	log.Printf("Starting server on port %s ", port)
	e.Logger.Fatal(e.Start(port))
}
