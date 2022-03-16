package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(с Config) (*sqlx.DB, error) {
	confString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		с.Host, с.Port, с.Username, с.DBName, с.Password, с.SSLMode,
	)

	db, err := sqlx.Open("postgres", confString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return db, nil
}
