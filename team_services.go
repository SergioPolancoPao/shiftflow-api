package main

import (
	"fmt"

	"gorm.io/gorm"
)

type ITeamRepository interface {
	CreateTeam(*Team) (tx *gorm.DB)
	GetTeams(GetTeamsQueryParams, *[]Team) (tx *gorm.DB)
	GetTeam(string, *Team) (tx *gorm.DB)
	DeleteTeam(*Team) (tx *gorm.DB)
}

type TeamService struct {
	tRepository ITeamRepository
}

func (ts *TeamService) CreateTeam(team *Team) (tx *gorm.DB) {
	return ts.tRepository.CreateTeam(team)
}

func (ts *TeamService) GetTeams(filter GetTeamsQueryParams, teams *[]Team) (error) {
	if err := ts.tRepository.GetTeams(filter, teams).Error; err != nil {
		return fmt.Errorf("error retrieving team list: %w", err)
	}

	return nil
}

func (ts *TeamService) GetTeam(id string, team *Team) (error) {
	if err := ts.tRepository.GetTeam(id, team).Error; err != nil {
		return fmt.Errorf("error checking if team exists: %w", err)
	}

	return nil
}

func (ts *TeamService) DeleteTeam(id string, team *Team) (error) {
	err := ts.GetTeam(id, team)

	if err != nil {
		return fmt.Errorf("error checking if team exists: %w", err)
	}

	if err := ts.tRepository.DeleteTeam(team).Error; err != nil {
		return fmt.Errorf("error deleting team: %w", err)
	}

	return nil
}

func NewTeamService(tr ITeamRepository) *TeamService {
	return &TeamService{
		tRepository: tr,
	}
}
