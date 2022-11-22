package repository

import (
	"fmt"
	models "github.com/ivanovamir/grpc-besm/api/models/srv_1c"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(fmt.Sprintf(fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.DBName))), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		//Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	} else {
		if err := db.AutoMigrate(&models.Product1c{}); err != nil {
			return nil, err
		} else {
			return db, nil
		}
	}
}
