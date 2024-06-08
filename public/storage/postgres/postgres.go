package postgres

import (
	"public/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db      *sql.DB
	Parties storage.Parties
	Publics storage.Publics
}

const (
	host     = "localhost"
	password = "1234"
	user     = "postgres"
	port     = 5432
	dbname   = "newdata"
)

func DBConnect() (*Storage, error) {
	psql := fmt.Sprintf("host=%s password=%s user=%s port=%d dbname=%s", host, password, user, port, dbname)

	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	pr := NewPartyRepo(db)
	prr := NewPublicRepo(db)

	storage := &Storage{db: db, Parties: pr, Publics: prr}

	return storage, nil
}

func (stg *Storage) Party() storage.Parties {
	if stg.Parties == nil {
		stg.Parties = NewPartyRepo(stg.db)
	}
	return stg.Parties
}

func (stg *Storage) Public() storage.Publics {
	if stg.Publics == nil {
		stg.Publics = NewPublicRepo(stg.db)
	}
	return stg.Publics
}
