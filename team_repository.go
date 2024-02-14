package main

import (
	"gorm.io/gorm"
)

type ITeamRepository interface {
	CreateTeam(*Team) (tx *gorm.DB)
	GetTeams(GetTeamsQueryParams, *[]Team) (tx *gorm.DB)
	GetTeam(string, *Team) (tx *gorm.DB)
	DeleteTeam(*Team) (tx *gorm.DB)
}

type TeamRepository struct {
	dbClient *gorm.DB
}

func (tr *TeamRepository) CreateTeam(team *Team) (tx *gorm.DB) {
	return tr.dbClient.Create(&team)
}

func (tr *TeamRepository) GetTeams(filter GetTeamsQueryParams, teams *[]Team) (tx *gorm.DB) {

	return tr.dbClient.Find(&teams, filter)
}

func (tr *TeamRepository) GetTeam(id string, team *Team) (tx *gorm.DB) {
	return tr.dbClient.Where("id = ?", id).First(&team)
}

func (tr *TeamRepository) DeleteTeam(team *Team) (tx *gorm.DB) {
	return tr.dbClient.Delete(&team)
}

func NewTeamRepository(db *gorm.DB) ITeamRepository {
	return &TeamRepository{
		dbClient: db,
	}
}
