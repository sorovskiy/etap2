package repository

import (
	"database/sql"
	"net/http"
)

type Repo struct {
	db         *sql.DB
	httpClient http.Client
}

func NewRepository(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}
