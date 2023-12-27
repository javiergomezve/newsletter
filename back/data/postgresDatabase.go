package data

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"newsletter-back/models"
)

type PostgresDatabase struct {
	DB       *gorm.DB
	Username string
	Password string
	Db       string
	Host     string
	Port     string
	SslMode  string
}

type PostgresDatabaseOptions struct {
	Username string
	Password string
	Db       string
	Host     string
	Port     string
	SslMode  string
}

func NewPostgresDatabase(options PostgresDatabaseOptions) (*gorm.DB, error) {
	p := PostgresDatabase{
		DB:       nil,
		Username: options.Username,
		Password: options.Password,
		Db:       options.Db,
		Host:     options.Host,
		Port:     options.Port,
		SslMode:  options.SslMode,
	}

	err := p.connect()
	if err != nil {
		return nil, err
	}

	return p.DB, nil
}

func (p *PostgresDatabase) connect() error {
	if p.DB != nil {
		return nil
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=%s",
		p.Username, p.Password, p.Db, p.Host, p.SslMode,
	)

	var err error
	p.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = p.DB.AutoMigrate(
		&models.User{},
		&models.List{},
		&models.Media{},
		&models.Recipient{},
		&models.Newsletter{},
	)
	if err != nil {
		return err
	}

	return nil
}
