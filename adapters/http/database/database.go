package database

import (
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type (
	Db interface {
		Connect() error
	}

	Database struct {
		Dsn  string
		Conn *gorm.DB
		Env  string
	}
)

func NewDatabase() *Database {
	return &Database{
		Dsn: "",
		Env: "development",
	}
}

func (d *Database) Connect() error {
	db, err := gorm.Open(
		mysql.Open(d.Dsn),
		&gorm.Config{},
	)

	sql, err := db.DB()
	if err != nil {
		return err
	}

	sql.SetMaxIdleConns(5)
	sql.SetConnMaxIdleTime(time.Second)

	sql.SetMaxOpenConns(5)
	sql.SetConnMaxLifetime(time.Second)

	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}

	d.Conn = db

	return err
}
