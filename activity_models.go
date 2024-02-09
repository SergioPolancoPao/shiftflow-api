package main

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Name   string
	TeamID uint
	Team   Team

	CurrentLeader int
	NextLeader    int
	ActivityID     uint
	Periodicity    string
}

type ActivityType struct {
	gorm.Model
	Name string
}
