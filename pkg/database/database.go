package database

import (
	"fmt"

	"github.com/HEEPOKE/go-gin-hexagonal-api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSsl, cfg.DB_TIMEZONE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectAndCloseDatabase(cfg *config.Config) *gorm.DB {
	db, err := ConnectDatabase(cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}()

	return db
}
