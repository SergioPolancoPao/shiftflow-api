package main

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var dbName string = "shiftflow"

func TestTeamService(t *testing.T) {
	setup()
	t.Run("should create team", func(t *testing.T) {
		team := &Team{
			Name: "team test 1",
		}

		tr := NewTeamRepository(db)

		res := tr.CreateTeam(team)

		assert.Nil(t, res.Error)
	})

	t.Run("should throw an error because name is duplicated", func(t *testing.T) {
		team := &Team{
			Name: "team test 1",
		}

		tr := NewTeamRepository(db)

		res := tr.CreateTeam(team)

		assert.NotNil(t, res.Error)
	})

	t.Run("should retrieve all teams", func(t *testing.T) {
		var teams []Team
		teams = append(teams, Team{Name: "team test 2"})
		teams = append(teams, Team{Name: "team test 3"})

		tr := NewTeamRepository(db)

		for _, team := range teams {
			tr.CreateTeam(&team)
		}

		f := &GetTeamsQueryParams{
			Name: "",
			ID:   0,
		}

		var rTeams []Team

		err := tr.GetTeams(*f, &rTeams).Error

		assert.Nil(t, err)
		assert.Equal(t, len(rTeams), 3)
	})

	t.Run("should filter teams by name", func(t *testing.T) {
		tr := NewTeamRepository(db)

		f := &GetTeamsQueryParams{
			Name: "team test 2",
			ID:   0,
		}

		var rTeams []Team

		err := tr.GetTeams(*f, &rTeams).Error

		assert.Nil(t, err)
		assert.Equal(t, len(rTeams), 1)
		assert.Equal(t, rTeams[0].Name, "team test 2")
	})

	t.Run("should retrieve a single team by id", func(t *testing.T) {
		tr := NewTeamRepository(db)

		f := &GetTeamsQueryParams{
			Name: "",
			ID:   0,
		}

		var teams []Team

		err := tr.GetTeams(*f, &teams).Error

		assert.Nil(t, err)

		var team Team

		tErr := tr.GetTeam(strconv.FormatUint(uint64(teams[0].ID), 10), &team).Error

		assert.Equal(t, team.Name, teams[0].Name)
		assert.Nil(t, tErr)
	})

	t.Run("should throw not found error looking for a single team by id", func(t *testing.T) {
		tr := NewTeamRepository(db)

		var team Team

		err := tr.GetTeam("1000", &team)

		assert.NotNil(t, err)
	})

	t.Run("should delete a team", func(t *testing.T) {
		tr := NewTeamRepository(db)

		f := &GetTeamsQueryParams{
			Name: "",
			ID:   0,
		}

		var teams []Team

		err := tr.GetTeams(*f, &teams).Error

		assert.Nil(t, err)

		team := teams[0]

		dErr := tr.DeleteTeam(&team)

		assert.NotNil(t, dErr)
	})

	teardownGlobal()
}

func setup() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: 1,
		},
	)

	var err error

	db, err = InitDB(newLogger)

	if err != nil {
		log.Fatal(err)
	}
}

func teardownGlobal() {
	connection, err := db.DB()

	if err != nil {
		log.Fatalln(err)
	}

	defer connection.Close()
}
