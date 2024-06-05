package postgres

import "database/sql"

type PublicRepo struct{
	db *sql.DB
}

func NewPublicRepo(db *sql.DB) *PublicRepo  {
	return &PublicRepo{db: db}
}