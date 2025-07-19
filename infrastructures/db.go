package infrastructures

import (
	"fmt"
	"log"
	"time"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.DatabaseName)

	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.DB.ConnMaxLifetimeInSeconds) * time.Second)

	return db, nil
}
