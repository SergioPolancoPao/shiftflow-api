package main

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name      string
	Teammates []Teammate `gorm:"many2many:team_teammate;"`
}

type Teammate struct {
	gorm.Model
	Name  string
	Email string
}

type TeamTeammate struct {
	TeamID     uint `gorm:"primaryKey"`
	TeammateID uint `gorm:"primaryKey"`
	position   int  `gorm:"primaryKey"`
	CreatedAt  time.Time
}