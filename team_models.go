package main

import (
	"time"

	"gorm.io/gorm"
)

type Status string

type Team struct {
	gorm.Model
	Name      string     `json:"name" gorm:"not null;index;unique"`
	Teammates []Teammate `gorm:"many2many:team_teammate;"`
}

type Teammate struct {
	gorm.Model
	Name  string `gorm:"not null;index"`
	Email string
}

type TeamTeammate struct {
	TeamID     uint `gorm:"primaryKey"`
	TeammateID uint `gorm:"primaryKey"`
	Position   int  `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
