package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type TeamRepositoryMock struct {
	CreateTeamCalls [] *Team
	CreateTeamReturn *gorm.DB
}

func (m *TeamRepositoryMock) CreateTeam(team *Team) (tx *gorm.DB) {
	m.CreateTeamCalls = append(m.CreateTeamCalls, team)
	return m.CreateTeamReturn
}

func (m *TeamRepositoryMock) GetTeams(filter GetTeamsQueryParams, teams *[]Team) (tx *gorm.DB) {
	return
}

func (m *TeamRepositoryMock) GetTeam(id string, team *Team) (tx *gorm.DB) {
	return
}

func (m *TeamRepositoryMock) DeleteTeam(team *Team) (tx *gorm.DB) {
	return
}

func NewTeamRepositoryMock() ITeamRepository {
	return &TeamRepositoryMock{}
}

func TestTeamService(t *testing.T) {
	t.Run("should create team", func(t *testing.T) {
		team := &Team{
			Name: "team test 1",
		}
		tr := new(TeamRepositoryMock)
		tr.CreateTeamReturn = &gorm.DB{
			Error: nil,
		}

		ts := NewTeamService(tr)
		err := ts.CreateTeam(team)

		assert.Equal(t, len(tr.CreateTeamCalls), 1)
		assert.Equal(t, tr.CreateTeamCalls[0], team)
		assert.Nil(t, err.Error)
	})
}
