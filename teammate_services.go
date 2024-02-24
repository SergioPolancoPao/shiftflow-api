package main

import (
	"fmt"

	"gorm.io/gorm"
)

type ITeammateRepository interface {
	CreateTeammate(*Teammate) (tx *gorm.DB)
	GetTeammates(GetTeammatesQueryParams, *[]Teammate) (tx *gorm.DB)
	GetTeammate(string, *Teammate) (tx *gorm.DB)
	DeleteTeammate(*Teammate) (tx *gorm.DB)
}

type TeammateService struct {
	tmRepository ITeammateRepository
}

func (tms *TeammateService) CreateTeammate(teammate *Teammate) (tx *gorm.DB) {
	return tms.tmRepository.CreateTeammate(teammate)
}

func (tms *TeammateService) GetTeammates(filter GetTeammatesQueryParams, teams *[]Teammate) error {
	if err := tms.tmRepository.GetTeammates(filter, teams).Error; err != nil {
		return fmt.Errorf("error retrieving teammate list: %w", err)
	}

	return nil
}

func (tms *TeammateService) GetTeammate(id string, teammate *Teammate) error {
	if err := tms.tmRepository.GetTeammate(id, teammate).Error; err != nil {
		return fmt.Errorf("error checking if teammate exists: %w", err)
	}

	return nil
}

func (tms *TeammateService) DeleteTeam(id string, teammate *Teammate) error {
	err := tms.GetTeammate(id, teammate)

	if err != nil {
		return fmt.Errorf("error checking if teammate exists: %w", err)
	}

	if err := tms.tmRepository.DeleteTeammate(teammate).Error; err != nil {
		return fmt.Errorf("error deleting teammate: %w", err)
	}

	return nil
}

func NewTeammateService(tmr ITeammateRepository) *TeammateService {
	return &TeammateService{
		tmRepository: tmr,
	}
}
