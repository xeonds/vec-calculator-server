package lib

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(config *DatabaseConfig, migrator func(*gorm.DB) error) *gorm.DB {
	var db *gorm.DB
	var err error
	switch config.Type {
	case "mysql":
		dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.DB), &gorm.Config{})
	}
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if config.Migrate {
		if migrator == nil {
			log.Fatalf("Migrator is nil")
		}
		if err = migrator(db); err != nil {
			log.Fatalf("Failed to migrate tables: %v", err)
		}
	}
	return db
}
