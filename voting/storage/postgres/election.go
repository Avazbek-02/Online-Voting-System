package postgres

import (
	"database/sql"
	"time"

	v "all/voting/genproto/genproto/voting"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ElectionRepo struct {
	db *sql.DB
}

func NewElectionRepo(db *sql.DB) *ElectionRepo {
	return &ElectionRepo{db: db}
}

func (repo *ElectionRepo) Create(election *v.CreateElectionRequest) (*v.ElectionResponse, error) {
	query := `INSERT INTO elections (name, date) VALUES ($1, $2) RETURNING id, name, date`
	var id string
	err := repo.db.QueryRow(query, election.Name, election.Date.AsTime()).Scan(&id, &election.Name, &election.Date)
	if err != nil {
		return nil, err
	}
	return &v.ElectionResponse{
		Id:   id,
		Name: election.Name,
		Date: election.Date,
	}, nil
}

func (repo *ElectionRepo) GetByID(id string) (*v.ElectionResponse, error) {
	query := "SELECT id, name, date FROM elections WHERE id = $1"
	row := repo.db.QueryRow(query, id)

	var election v.ElectionResponse
	var date time.Time
	err := row.Scan(&election.Id, &election.Name, &date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	election.Date = timestamppb.New(date)
	return &election, nil
}

func (repo *ElectionRepo) Update(id string, election *v.CreateElectionRequest) (*v.ElectionResponse, error) {
	query := `UPDATE elections SET name = $2, date = $3 WHERE id = $1 RETURNING id, name, date`
	var updatedElection v.ElectionResponse
	var date time.Time
	err := repo.db.QueryRow(query, id, election.Name, election.Date.AsTime()).Scan(&updatedElection.Id, &updatedElection.Name, &date)
	if err != nil {
		return nil, err
	}
	updatedElection.Date = timestamppb.New(date)
	return &updatedElection, nil
}


func (repo *ElectionRepo) Delete(id string) error {
	query := "DELETE FROM elections WHERE id = $1"
	_, err := repo.db.Exec(query, id)
	return err
}
