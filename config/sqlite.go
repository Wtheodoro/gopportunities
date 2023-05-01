package config

import (
	"os"

	"github.com/Wtheodoro/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	
	// DSN = data source name
	dsn := "./db/main.db"

	// Chack if the database file exists
	_, err := os.Stat(dsn)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		// create db file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dsn)
		if err != nil {
			return nil, err
		}

		// Always rememver to close the file
		file.Close()
	}

	// create DB and connect, DSN = data source name
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Sqlite opening error: %v", err)
		return nil, err
	}

	// migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("Sqlite automigration error: %v", err)
	}

	// return DB
	return db, nil
}