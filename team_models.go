package main

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name      string     `json:"name" gorm:"not null"`
	Teammates []Teammate `gorm:"many2many:team_teammate;"`
}

type Teammate struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string
}

type TeamTeammate struct {
	TeamID     uint `gorm:"primaryKey"`
	TeammateID uint `gorm:"primaryKey"`
	Position   int  `gorm:"primaryKey"`
	CreatedAt  time.Time
}
