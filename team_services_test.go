package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestTeamService(t *testing.T) {
	t.Run("should create team", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		team := &Team{
			Name: "team test 1",
		}
		tr := NewMockITeamRepository(ctrl)
		ts := NewTeamService(tr)

		response := &gorm.DB{
			Error: nil,
		}
		tr.EXPECT().CreateTeam(team).Return(response)
		r := ts.CreateTeam(team)

		assert.Nil(t, r.Error)
	})

	t.Run("should return an error in create team method", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		team := &Team{
			Name: "team test 1",
		}
		tr := NewMockITeamRepository(ctrl)
		ts := NewTeamService(tr)

		response := &gorm.DB{
			Error: gorm.ErrDuplicatedKey,
		}
		tr.EXPECT().CreateTeam(team).Return(response)
		err := ts.CreateTeam(team)

		assert.NotNil(t, err.Error)
		assert.True(t, errors.Is(gorm.ErrDuplicatedKey, err.Error))
	})

	t.Run("should call get teams properly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		filter := &GetTeamsQueryParams{
			Name: "team test 1",
		}
		tr := NewMockITeamRepository(ctrl)
		ts := NewTeamService(tr)

		response := &gorm.DB{
			Error: nil,
		}
		var teams []Team
		tr.EXPECT().GetTeams(*filter, &teams).Return(response)
		err := ts.GetTeams(*filter, &teams)

		assert.Nil(t, err)
	})

	t.Run("should call get teams with error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		filter := &GetTeamsQueryParams{
			Name: "team test 1",
		}
		tr := NewMockITeamRepository(ctrl)
		ts := NewTeamService(tr)

		response := &gorm.DB{
			Error: gorm.ErrInvalidField,
		}
		var teams []Team
		tr.EXPECT().GetTeams(*filter, &teams).Return(response)
		err := ts.GetTeams(*filter, &teams)

		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, gorm.ErrInvalidField))
	})

	t.Run("should call get team properly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "1"
		tr := NewMockITeamRepository(ctrl)
		ts := NewTeamService(tr)

		response := &gorm.DB{
			Error: nil,
		}
		var team Team
		tr.EXPECT().GetTeam(id, &team).Return(response)
		err := ts.GetTeam(id, &team)

		assert.Nil(t, err)
	})

	t.Run("should call get team with error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "1"
		tr := NewMockITeamRepository(ctrl)
		ts := NewTeamService(tr)

		response := &gorm.DB{
			Error: gorm.ErrRecordNotFound,
		}
		var team Team
		tr.EXPECT().GetTeam(id, &team).Return(response)
		err := ts.GetTeam(id, &team)

		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	})

	t.Run("should call delete team properly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tr := NewMockITeamRepository(ctrl)
		ts := NewTeamService(tr)
		gResponse := &gorm.DB{
			Error: nil,
		}
		dResponse := &gorm.DB{
			Error: nil,
		}
		var team *Team
		id := "1"
		tr.EXPECT().GetTeam(id, team).Return(gResponse)
		tr.EXPECT().DeleteTeam(team).Return(dResponse)
		err := ts.DeleteTeam(id, team)

		assert.Nil(t, err)
	})

	t.Run("should call delete team with not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tr := NewMockITeamRepository(ctrl)
		ts := NewTeamService(tr)
		gResponse := &gorm.DB{
			Error: gorm.ErrRecordNotFound,
		}
		var team *Team
		id := "1"
		tr.EXPECT().GetTeam(id, team).Return(gResponse)
		err := ts.DeleteTeam(id, team)

		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	})

	t.Run("should call delete team with delete error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tr := NewMockITeamRepository(ctrl)
		ts := NewTeamService(tr)
		gResponse := &gorm.DB{
			Error: nil,
		}
		dResponse := &gorm.DB{
			Error: gorm.ErrInvalidField,
		}
		var team *Team
		id := "1"
		tr.EXPECT().GetTeam(id, team).Return(gResponse)
		tr.EXPECT().DeleteTeam(team).Return(dResponse)
		err := ts.DeleteTeam(id, team)

		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, gorm.ErrInvalidField))
	})
}
