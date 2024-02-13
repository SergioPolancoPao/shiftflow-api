package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func router(server *echo.Echo) error {
	server.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	server.GET("/metrics", echoprometheus.NewHandler())

	return nil
}

func InitDB(logger logger.Interface) (*gorm.DB, error) {
	db, err := NewDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
		logger,
	)

	if err != nil {
		return nil, fmt.Errorf("error initializing db: %w", err)
	}

	setupTableErr := db.SetupJoinTable(&Team{}, "Teammates", &TeamTeammate{})

	if setupTableErr != nil {
		return nil, setupTableErr
	}

	mErr := db.AutoMigrate(&Team{}, &Teammate{}, &ActivityType{}, &Activity{})

	if mErr != nil {
		return nil, fmt.Errorf("error migrating db: %w", mErr)
	}

	return db, nil
}

func getLogLevel(logLevelStr string) logger.LogLevel {
	var logLevel logger.LogLevel
	switch logLevelStr {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Info
	}

	return logLevel
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logLevelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))

	logLevel := getLogLevel(logLevelStr)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logLevel,
		},
	)

	db, err := InitDB(newLogger)

	if err != nil {
		panic(err)
	}

	log.Print("Database server is running")

	e := echo.New()

	e.Use(echoprometheus.NewMiddleware("shiftflow"))
	e.Use(middleware.Logger())

	validator := validator.New()

	e.Validator = &CustomValidator{validator}

	if rErr := router(e); rErr != nil {
		panic(err)
	}

	tr := NewTeamRepository(db)
	ts := NewTeamService(tr)
	NewTeamController(e, ts)

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
