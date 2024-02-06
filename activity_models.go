package main

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Name   string
	TeamID uint
	Team   Team

	current_leader int
	next_leader    int
	ActivityID     uint
	Periodicity    string
}

type ActivityType struct {
	gorm.Model
	Name string
}
