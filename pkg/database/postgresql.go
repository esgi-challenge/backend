package database

import (
	"fmt"

	"github.com/esgi-challenge/backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient(c *config.Config) (*gorm.DB, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.Dbname,
	)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
