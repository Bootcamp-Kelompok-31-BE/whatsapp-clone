package infrastructures

import (
	"fmt"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	// "time"
)

type PostgresDatabase struct {
	db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *PostgresDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(
			fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s timezone=%s",
				conf.Db.Host,
				conf.Db.User,
				conf.Db.Password,
				conf.Db.DbName,
				conf.Db.Port,
				conf.Db.SslMode,
				conf.Db.TimeZone,
			),
		),
			&gorm.Config{},
		)

		if err != nil {
			panic(err)
		}

		dbInstance = &PostgresDatabase{
			db: db,
		}
	})

	return dbInstance
}

func (p *PostgresDatabase) GetInstance() *gorm.DB {
	return dbInstance.db
}


func NewDB(config *config.Config) (*gorm.DB, error)  {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Db.Host, config.Db.Port, config.Db.User, config.Db.Password, config.Db.DbName)

	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	/*
	sqlDb, err := db.Db()
	if err != nil {
		return nil, err
	}
	
	sqlDb.SetMaxIdleConns(config.Db.MaxIdleConns)
	sqlDb.SetMaxOpenConns(config.Db.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(config.Db.ConnMaxLifetimeInSeconds) * time.Second) */

	return db, nil
}