package main

import "gorm.io/gorm"

type ITeamService interface {
	CreateTeam() string
}

type TeamService struct {
	dbClient *gorm.DB
}

func (ts *TeamService) CreateTeam() string {
	return "Success save"
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{
		dbClient: db,
	}
}
