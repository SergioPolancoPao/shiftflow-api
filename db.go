package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(host, port, username, password, dbname, sslmode string, logger logger.Interface) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, username, password, dbname, sslmode,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	return db, nil
}
