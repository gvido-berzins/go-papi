package storage

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupDB sets up an Sqlite database and returns the ORM model.
func SetupDB() *gorm.DB {
	log.Trace().Msg("setting up database")
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&Command{})
	log.Trace().Msg("database setup done")
	return db
}
