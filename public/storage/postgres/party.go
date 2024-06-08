package postgres

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"

	v "public/genproto/genproto/public"
)

type PartyRepo struct {
	db *sql.DB
}

func NewPartyRepo(db *sql.DB) *PartyRepo {
	return &PartyRepo{db: db}
}

func (pt *PartyRepo) CreateParty(req *v.CreatePartyRequest) (*v.PartyResponse, error) {
	query := `
		INSERT INTO party(id, name, slogan, opened_at, description) VALUES($1, $2, $3, $4, $5)
	`
	id := uuid.NewV4().String()

	_, err := pt.db.Exec(query, id, req.Name, req.Slogan, req.OpenedAt, req.Description)
	if err != nil {
		return nil, err
	}

	return &v.PartyResponse{
		Id:          id,
		Name:        req.Name,
		Slogan:      req.Slogan,
		OpenedAt:    req.OpenedAt,
		Description: req.Description,
	}, nil
}

func (pt *PartyRepo) SelectParty(id *v.GetPartyInfoRequest) (*v.PartyResponse, error) {
	query := `
		SELECT id, name, slogan, opened_at, description FROM party WHERE id = $1
	`
	var party v.PartyResponse

	err := pt.db.QueryRow(query, id.Id).Scan(
		&party.Id,
		&party.Name,
		&party.Slogan,
		&party.OpenedAt,
		&party.Description,
	)
	if err != nil {
		return nil, err
	}

	return &party, nil
}

func (pt *PartyRepo) UpdateParty(req *v.UpdatePartyRequest) (*v.PartyResponse, error) {
	query := `
		UPDATE party SET name = $2, slogan = $3, opened_at = $4, description = $5 WHERE id = $1
	`
	_, err := pt.db.Exec(query, req.Id, req.Name, req.Slogan, req.OpenedAt, req.Description)
	if err != nil {
		return nil, err
	}

	return pt.SelectParty(&v.GetPartyInfoRequest{Id: req.Id})
}

func (pt *PartyRepo) DeleteParty(id *v.DeletePartyRequest) (*v.Voidd, error) {
	query := `
		DELETE FROM party WHERE id = $1
	`
	_, err := pt.db.Exec(query, id)
	if err != nil {
		return nil, err
	}

	return &v.Voidd{}, nil
}
