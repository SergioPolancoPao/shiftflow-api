package main

import (
	"fmt"

	"gorm.io/gorm"
)

type TeamService struct {
	tRepository ITeamRepository
}

func (ts *TeamService) CreateTeam(team *Team) (tx *gorm.DB) {
	return ts.tRepository.CreateTeam(team)
}

func (ts *TeamService) GetTeams(filter GetTeamsQueryParams) ([]Team, error) {
	var teams []Team

	if err := ts.tRepository.GetTeams(filter, &teams).Error; err != nil {
		return nil, fmt.Errorf("error retrieving team list: %w", err)
	}

	return teams, nil
}

func (ts *TeamService) GetTeam(id string) (*Team, error) {
	var team Team

	if err := ts.tRepository.GetTeam(id, &team).Error; err != nil {
		return nil, fmt.Errorf("error checking if team exists: %w", err)
	}

	return &team, nil
}

func (ts *TeamService) DeleteTeam(id string) (*Team, error) {
	team, err := ts.GetTeam(id)

	if err != nil {
		return nil, fmt.Errorf("error checking if team exists: %w", err)
	}

	if err := ts.tRepository.DeleteTeam(team).Error; err != nil {
		return nil, fmt.Errorf("error deleting team: %w", err)
	}

	return team, nil
}

func NewTeamService(tr ITeamRepository) *TeamService {
	return &TeamService{
		tRepository: tr,
	}
}
