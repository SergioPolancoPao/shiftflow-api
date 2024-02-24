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

var dbTeamRepository *gorm.DB

func TestTeamRepository(t *testing.T) {
	setupTeamRepositoryTest()
	t.Run("should create team", func(t *testing.T) {
		team := &Team{
			Name: "team test 1",
		}

		tr := NewTeamRepository(dbTeamRepository)

		res := tr.CreateTeam(team)

		assert.Nil(t, res.Error)
	})

	t.Run("should throw an error because name is duplicated", func(t *testing.T) {
		team := &Team{
			Name: "team test 1",
		}

		tr := NewTeamRepository(dbTeamRepository)

		res := tr.CreateTeam(team)

		assert.NotNil(t, res.Error)
	})

	t.Run("should retrieve all teams", func(t *testing.T) {
		var teams []Team
		teams = append(teams, Team{Name: "team test 2"})
		teams = append(teams, Team{Name: "team test 3"})

		tr := NewTeamRepository(dbTeamRepository)

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
		tr := NewTeamRepository(dbTeamRepository)

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
		tr := NewTeamRepository(dbTeamRepository)

		f := &GetTeamsQueryParams{
			Name: "",
			ID:   0,
		}

		var teams []Team

		err := tr.GetTeams(*f, &teams).Error

		assert.Nil(t, err)

		var team Team

		tErr := tr.GetTeam(strconv.FormatUint(uint64(teams[0].ID), 10), &team).Error

		assert.Nil(t, tErr)
		assert.Equal(t, team.Name, teams[0].Name)
	})

	t.Run("should throw not found error looking for a single team by id", func(t *testing.T) {
		tr := NewTeamRepository(dbTeamRepository)

		var team Team

		err := tr.GetTeam("1000", &team)

		assert.NotNil(t, err)
	})

	t.Run("should delete a team", func(t *testing.T) {
		tr := NewTeamRepository(dbTeamRepository)

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

	teardownGlobalTeamRepositoryTest()
}

func setupTeamRepositoryTest() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: 1,
		},
	)

	var err error

	dbTeamRepository, err = InitDB(newLogger)
	if err != nil {
		log.Fatalf("Error initializing db %s", err)
	}
}

func teardownGlobalTeamRepositoryTest() {
	connection, err := dbTeamRepository.DB()
	if err != nil {
		log.Fatalln(err)
	}

	defer connection.Close()
}
