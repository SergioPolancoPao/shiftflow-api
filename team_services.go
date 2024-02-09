package main

import (
	"fmt"

	"gorm.io/gorm"
)

type TeamService struct {
	dbClient *gorm.DB
}

func (ts *TeamService) CreateTeam(team *Team) (tx *gorm.DB) {
	return ts.dbClient.Create(&team)
}

func (ts *TeamService) GetTeams(name string) ([]Team, error) {
	var teams []Team
	query := ts.dbClient.Model(&Team{})

	if name != "" {
		query = query.Where("name = ?", name)
	}
	
	if err := query.Find(&teams).Error; err != nil {
		return nil, fmt.Errorf("error retrieving team list: %w", err)
	}

	return teams, nil
}

func (ts *TeamService) GetTeam(id string) (*Team, error) {
	var team Team

	if err := ts.dbClient.Where("id = ?", id).First(&team).Error; err != nil {
		return nil, fmt.Errorf("error checking if team exists: %w", err)
	}

	return &team, nil
}

func (ts *TeamService) DeleteTeam(id string) (*Team, error) {
	team, err := ts.GetTeam(id)

	if err != nil {
		return nil, fmt.Errorf("error checking if team exists: %w", err)
	}

	if err := ts.dbClient.Delete(&team).Error; err != nil {
		return nil, fmt.Errorf("error deleting team: %w", err)
	}

	return team, nil
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{
		dbClient: db,
	}
}
