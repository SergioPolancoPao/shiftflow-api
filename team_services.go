package main

import "gorm.io/gorm"

type ITeamService interface {
	CreateTeam() string
}

type TeamService struct {
	dbClient *gorm.DB
}

func (ts *TeamService) CreateTeam(team *Team) (tx *gorm.DB) {
	return ts.dbClient.Create(&team)
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{
		dbClient: db,
	}
}
