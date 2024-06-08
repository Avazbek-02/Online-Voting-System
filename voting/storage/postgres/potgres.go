package postgres

import (
	"database/sql"
	"fmt"

	"all/voting/storage"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	password = "1234"
	user     = "postgres"
	port     = 5432
	dbname   = "newdata"
)

type Storage struct {
	db          *sql.DB
	Candidates  storage.Candidates
	Elections   storage.Elections
	PublicVotes storage.PublicVotes
}

func DBConnect() (*Storage, error) {
	psql := fmt.Sprintf("host=%s password=%s user=%s port=%d dbname=%s sslmode=disable", host, password, user, port, dbname)

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
	pv := NewPublicVoteRepo(db)

	return &Storage{db: db, Candidates: cr, Elections: er, PublicVotes: pv}, nil
}

func (stg *Storage) Candidate() storage.Candidates {
	if stg.Candidates == nil {
		stg.Candidates = NewCandidateRepo(stg.db)
	}
	return stg.Candidates
}

func (stg *Storage) Election() storage.Elections {
	if stg.Elections == nil {
		stg.Elections = NewElectionRepo(stg.db)
	}
	return stg.Elections
}

func (stg *Storage) PublicVote() storage.PublicVotes {
	if stg.PublicVotes == nil {
		stg.PublicVotes = NewPublicVoteRepo(stg.db)
	}
	return stg.PublicVotes
}
