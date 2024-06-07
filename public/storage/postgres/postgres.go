package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db         *sql.DB
	PartyRepo  *PartyRepo
	PublicRepo *PublicRepo
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

	return &Storage{db: db, PartyRepo: pr, PublicRepo: prr}, nil

}
