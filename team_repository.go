package main

import (
	"gorm.io/gorm"
)

type TeamRepository struct {
	dbClient *gorm.DB
}

func (tr *TeamRepository) CreateTeam(team *Team) (tx *gorm.DB) {
	return tr.dbClient.Create(&team)
}

func (tr *TeamRepository) GetTeams(filter GetTeamsQueryParams, teams *[]Team) (tx *gorm.DB) {
	return tr.dbClient.Scopes(CommonFields(filter)).Find(&teams)
}

func (tr *TeamRepository) GetTeam(id string, team *Team) (tx *gorm.DB) {
	return tr.dbClient.Where("id = ?", id).First(&team)
}

func (tr *TeamRepository) DeleteTeam(team *Team) (tx *gorm.DB) {
	return tr.dbClient.Delete(&team)
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		dbClient: db,
	}
}
