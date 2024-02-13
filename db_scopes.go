package main

import "gorm.io/gorm"

// TODO: Use an interface to this function
func CommonFields(filter GetTeamsQueryParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.ID != 0 {
			db = db.Where("id = ?", filter.ID)
		}

		if filter.Name != "" {
			db = db.Where("name = ?", filter.Name)
		}

		return db
	}
}
