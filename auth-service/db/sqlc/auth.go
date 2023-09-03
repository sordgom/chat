package sqlc

import (
	"database/sql"
)

type Auth interface {
	Querier
}

type SQLAuth struct {
	*Queries
	db *sql.DB
}

func NewAuth(db *sql.DB) Auth {
	return &SQLAuth{
		Queries: New(db),
		db:      db,
	}
}
