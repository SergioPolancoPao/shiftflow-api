package main

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Name          string `gorm:"not null;index"`
	TeamID        uint   `gorm:"not null;index"`
	Team          Team
	CurrentLeader int
	NextLeader    int
	ActivityID    uint `gorm:"not null;index"`
	Periodicity   string
}

type ActivityType struct {
	gorm.Model
	Name string `gorm:"not null;index;unique"`
}
