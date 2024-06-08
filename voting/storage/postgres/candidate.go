package postgres

import (
	"database/sql"

	v "all/voting/genproto/genproto/voting"

	uuid "github.com/satori/go.uuid"
)

type CandidateRepo struct {
	db *sql.DB
}

func NewCandidateRepo(db *sql.DB) *CandidateRepo {
	return &CandidateRepo{db: db}
}

func (repo *CandidateRepo) CreateCandidate(req *v.CreateCandidateRequest) (*v.CandidateResponse, error) {
	query := `
		INSERT INTO candidate(id, election_id, public_id, party_id) VALUES($1, $2, $3, $4)
	`
	id := uuid.NewV4().String()

	_, err := repo.db.Exec(query, id, req.ElectionId, req.PublicId, req.PartyId)
	if err != nil {
		return nil, err
	}

	return &v.CandidateResponse{
		Id:         id,
		ElectionId: req.ElectionId,
		PublicId:   req.PublicId,
		PartyId:    req.PartyId,
	}, nil
}

func (repo *CandidateRepo) GetCandidateInfo(req *v.GetCandidateInfoRequest) (*v.CandidateResponse, error) {
	query := `
		SELECT id, election_id, public_id, party_id FROM candidate WHERE id = $1
	`
	var candidate v.CandidateResponse
	err := repo.db.QueryRow(query, req.Id).Scan(
		&candidate.Id,
		&candidate.ElectionId,
		&candidate.PublicId,
		&candidate.PartyId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &candidate, nil
}

func (repo *CandidateRepo) UpdateCandidate(req *v.UpdateCandidateRequest) (*v.CandidateResponse, error) {
	query := `
		UPDATE candidate SET election_id = $2, public_id = $3, party_id = $4 WHERE id = $1
	`
	_, err := repo.db.Exec(query, req.Id, req.ElectionId, req.PublicId, req.PartyId)
	if err != nil {
		return nil, err
	}

	return repo.GetCandidateInfo(&v.GetCandidateInfoRequest{Id: req.Id})
}

func (repo *CandidateRepo) DeleteCandidate(req *v.DeleteCandidateRequest) (*v.Void1, error) {
	query := `
		DELETE FROM candidate WHERE id = $1
	`
	_, err := repo.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &v.Void1{}, nil
}
