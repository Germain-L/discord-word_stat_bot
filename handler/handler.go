package handler

import "database/sql"

type handler struct {
	Db *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}
