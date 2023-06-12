package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Cfg.DBHost, config.Cfg.DBUser, config.Cfg.DBPassword, config.Cfg.DBName, config.Cfg.DBPort, config.Cfg.DBSsl, config.Cfg.DB_TIMEZONE)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}

func ConnectAndCloseDatabase() *gorm.DB {
	db, err := ConnectDatabase()
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
