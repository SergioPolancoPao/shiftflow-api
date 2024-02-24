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

func router(server *echo.Echo, tc *TeamController, tmc *TeammateController) error {
	server.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	server.GET("/metrics", echoprometheus.NewHandler())

	// Team group
	tg := server.Group("/teams")
	tg.POST("", tc.CreateTeam)
	tg.GET("", tc.GetTeams)
	tg.GET("/:id", tc.GetTeam)
	tg.DELETE("/:id", tc.DeleteTeam)

	// Teammate group
	tmg := server.Group("/teammates")
	tmg.POST("", tmc.CreateTeammate)
	tmg.GET("", tmc.GetTeammates)
	tmg.GET("/:id", tmc.GetTeammate)
	tmg.DELETE("/:id", tmc.DeleteTeammate)

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

	if err := db.SetupJoinTable(&Team{}, "Teammates", &TeamTeammate{}); err != nil {
		return nil, fmt.Errorf("error setting up team_teammate table: %w", err)
	}

	if err := db.AutoMigrate(&Team{}, &Teammate{}, &ActivityType{}, &Activity{}); err != nil {
		return nil, fmt.Errorf("error migrating db: %w", err)
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
		log.Fatalf("Error loading .env file: %s", err)
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
		log.Fatalf("Error initializing db %s", err)
	}

	log.Print("Database server is running")

	e := echo.New()

	e.Use(echoprometheus.NewMiddleware("shiftflow"))
	e.Use(middleware.Logger())

	validator := validator.New()

	e.Validator = &CustomValidator{validator}

	// team
	tr := NewTeamRepository(db)
	ts := NewTeamService(tr)
	tc := NewTeamController(ts)

	//teammate
	tmr := NewTeammateRepository(db)
	tms := NewTeammateService(tmr)
	tmc := NewTeammateController(tms)

	if err := router(e, tc, tmc); err != nil {
		log.Fatalf("Error initializing router %s", err)
	}

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
