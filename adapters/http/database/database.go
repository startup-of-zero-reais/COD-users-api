package database

import (
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type (
	Db interface {
		Connect() error
	}

	Database struct {
		Dsn               string
		Conn              *gorm.DB
		Env               string
		ShouldAutoMigrate string
	}
)

func NewDatabase() *Database {
	return &Database{
		Dsn:               getEnv("MYSQL_COD_DSN", ""),
		Env:               getEnv("APPLICATION_ENV", "development"),
		ShouldAutoMigrate: getEnv("MUST_AUTOMIGRATE", "false"),
	}
}

func (d *Database) Connect() error {
	db, err := gorm.Open(
		mysql.Open(d.Dsn),
		&gorm.Config{},
	)

	_, err = db.DB()
	if err != nil {
		return err
	}

	if d.ShouldAutoMigrate == "true" {
		err = db.AutoMigrate(&entities.User{}, &entities.RecoverToken{})
		if err != nil {
			return err
		}
	}

	d.Conn = db

	return err
}

func getEnv(key, _default string) string {
	if envVar := os.Getenv(key); envVar != "" {
		return envVar
	}

	return _default
}
