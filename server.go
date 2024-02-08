package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func router(server *echo.Echo) error {
	server.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	return nil
}

func initDB() *gorm.DB {
	db, err := NewDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	if err != nil {
		panic(err)
	}

	db.SetupJoinTable(&Team{}, "Teammates", &TeamTeammate{})
	mErr := db.AutoMigrate(&Team{}, &Teammate{}, &ActivityType{}, &Activity{})

	if mErr != nil {
		panic(err)
	}

	return db
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := initDB()

	log.Print("Database server is running")

	e := echo.New()

	validator := validator.New()

	e.Validator = &CustomValidator{validator}

	router(e)

	ts := NewTeamService(db)
	NewTeamController(e, ts)

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
