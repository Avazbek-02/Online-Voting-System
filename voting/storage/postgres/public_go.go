package postgres

import (
	"database/sql"

	v "voting/genproto/genproto/voting"

	uuid "github.com/satori/go.uuid"
)

type PublicVoteRepo struct {
	db *sql.DB
}

func NewPublicVoteRepo(db *sql.DB) *PublicVoteRepo {
	return &PublicVoteRepo{db: db}
}

func (repo *PublicVoteRepo) CreatePublicVote(req *v.CreatePublicVoteRequest) (*v.PublicVoteResponse, error) {
	query := `
		INSERT INTO public_vote(id, election_id, public_id) VALUES($1, $2, $3)
	`
	id := uuid.NewV4().String()

	_, err := repo.db.Exec(query, id, req.ElectionId, req.PublicId)
	if err != nil {
		return nil, err
	}

	return &v.PublicVoteResponse{
		Id:         id,
		ElectionId: req.ElectionId,
		PublicId:   req.PublicId,
	}, nil
}

func (repo *PublicVoteRepo) GetPublicVoteInfo(req *v.GetPublicVoteInfoRequest) (*v.PublicVoteResponse, error) {
	query := `
		SELECT id, election_id, public_id FROM public_vote WHERE id = $1
	`
	var publicVote v.PublicVoteResponse
	err := repo.db.QueryRow(query, req.Id).Scan(
		&publicVote.Id,
		&publicVote.ElectionId,
		&publicVote.PublicId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &publicVote, nil
}

func (repo *PublicVoteRepo) UpdatePublicVote(req *v.UpdatePublicVoteRequest) (*v.PublicVoteResponse, error) {
	query := `
		UPDATE public_vote SET election_id = $2, public_id = $3 WHERE id = $1
	`
	_, err := repo.db.Exec(query, req.Id, req.ElectionId, req.PublicId)
	if err != nil {
		return nil, err
	}

	return &v.PublicVoteResponse{
		Id:         req.Id,
		ElectionId: req.ElectionId,
		PublicId:   req.PublicId,
	}, nil
}

func (repo *PublicVoteRepo) DeletePublicVote(req *v.DeletePublicVoteRequest) (*v.Void3, error) {
	query := `
		DELETE FROM public_vote WHERE id = $1
	`
	_, err := repo.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &v.Void3{}, nil
}
