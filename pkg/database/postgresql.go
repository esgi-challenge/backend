package database

import (
	"fmt"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient(c *config.Config) (*gorm.DB, error) {
	dbUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		c.Postgres.Host,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Dbname,
		c.Postgres.Port,
	)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := migrateDatabase(db); err != nil {
		return nil, err
	}

	return db, err
}

func migrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Example{},
		&models.User{},
		&models.School{},
		&models.Campus{},
		&models.Path{},
		&models.Class{},
		&models.Course{},
		&models.Schedule{},
	)

	if err != nil {
		return err
	}

	return nil
}
