package dbrepo

import (
	"database/sql"

	"github.com/penhele/bookings/internal/config"
	"github.com/penhele/bookings/internal/repository"
)

type postgreDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgreRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgreDBRepo {
		App: a,
		DB: conn,
	}
}
