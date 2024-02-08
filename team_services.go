package main

import (
	"gorm.io/gorm"
)

type ITeamService interface {
	CreateTeam() string
}

type TeamService struct {
	dbClient *gorm.DB
}

func (ts *TeamService) CreateTeam(team *Team) (tx *gorm.DB) {
	return ts.dbClient.Create(&team)
}

func (ts *TeamService) GetTeams(name string) ([]Team, error) {
	var teams []Team

	if res := ts.dbClient.Where("name = ?", name).Find(&teams); res.Error != nil {
		return nil, res.Error
	}

	return teams, nil
}

func (ts *TeamService) GetTeam(id string) (*Team, error) {
	var team Team

	if res := ts.dbClient.Where("id = ?", id).First(&team); res.Error != nil {
		return nil, res.Error
	}

	return &team, nil
}

func (ts *TeamService) DeleteTeam(id string) (*Team, error) {
	team, err := ts.GetTeam(id)

	if err != nil {
		return nil, err
	}

	if res := ts.dbClient.Delete(&team); res.Error != nil {
		return nil, res.Error
	}

	return team, nil
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{
		dbClient: db,
	}
}
