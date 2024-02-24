package main

import (
	"gorm.io/gorm"
)

type TeammateRepository struct {
	dbClient *gorm.DB
}

func (tr *TeammateRepository) CreateTeammate(teammate *Teammate) (tx *gorm.DB) {
	return tr.dbClient.Create(&teammate)
}

func (tr *TeammateRepository) GetTeammates(filter GetTeammatesQueryParams, teammates *[]Teammate) (tx *gorm.DB) {
	return tr.dbClient.Find(&teammates, filter)
}

func (tr *TeammateRepository) GetTeammate(id string, teammate *Teammate) (tx *gorm.DB) {
	return tr.dbClient.Where("id = ?", id).First(&teammate)
}

func (tr *TeammateRepository) DeleteTeammate(teammate *Teammate) (tx *gorm.DB) {
	return tr.dbClient.Delete(&teammate)
}

func NewTeammateRepository(db *gorm.DB) *TeammateRepository {
	return &TeammateRepository{
		dbClient: db,
	}
}
