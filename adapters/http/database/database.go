package database

import (
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	// Db é a 'interface' necessária para ter a estrutura de conexão com o banco de dados
	Db interface {
		Connect() error
	}

	// Database é a estrutura que administra a conexão e Env de bancos de dados
	Database struct {
		Dsn               string
		Conn              *gorm.DB
		Env               string
		ShouldAutoMigrate string
	}
)

// NewDatabase método construtor de Database
func NewDatabase() *Database {
	return &Database{
		Dsn:               utilities.GetEnv("MYSQL_COD_DSN", ""),
		Env:               utilities.GetEnv("APPLICATION_ENV", "development"),
		ShouldAutoMigrate: utilities.GetEnv("MUST_AUTOMIGRATE", "false"),
	}
}

// Connect é o método que efetua de fato a conexão com o banco de dados
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
