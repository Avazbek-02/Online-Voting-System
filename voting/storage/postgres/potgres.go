package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db             *sql.DB
	PublicVoteRepo *PublicVoteRepo
	Election       *ElectionRepo
	Candidate      *CandidateRepo
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

	cr := NewCandidateRepo(db)
	er := NewElectionRepo(db)
	pr := NewPublicVoteRepo(db)

	return &Storage{db: db, PublicVoteRepo: pr, Candidate: cr, Election: er}, nil

}
