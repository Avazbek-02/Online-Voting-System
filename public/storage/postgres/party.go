package postgres

import (
	"database/sql"

	"all/public/genproto/genproto"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PartyRepo struct {
	db *sql.DB
}

func NewPartyRepo(db *sql.DB) *PartyRepo {
	return &PartyRepo{db: db}
}

func (pt *PartyRepo) CreateParty(req *genproto.CreatePartyRequest) (*genproto.PartyResponse, error) {
	query := `
		INSERT INTO party(id, name, slogan, opened_at, description) VALUES($1, $2, $3, $4, $5)
	`
	id := uuid.NewV4().String()

	_, err := pt.db.Exec(query, id, req.Name, req.Slogan, req.OpenedAt.AsTime(), req.Description)
	if err != nil {
		return nil, err
	}

	return &genproto.PartyResponse{
		Id:          id,
		Name:        req.Name,
		Slogan:      req.Slogan,
		OpenedAt:    req.OpenedAt,
		Description: req.Description,
	}, nil
}

func (pt *PartyRepo) SelectParty(id *genproto.GetPartyInfoRequest) (*genproto.PartyResponse, error) {
	query := `
		SELECT id, name, slogan, opened_at, description FROM party WHERE id = $1
	`
	var party genproto.PartyResponse
	var openedAt sql.NullTime
	err := pt.db.QueryRow(query, id.Id).Scan(
		&party.Id,
		&party.Name,
		&party.Slogan,
		&openedAt,
		&party.Description,
	)
	if err != nil {
		return nil, err
	}

	if openedAt.Valid {
		party.OpenedAt = timestamppb.New(openedAt.Time)
	}

	return &party, nil
}

func (pt *PartyRepo) UpdateParty(id string, req *genproto.CreatePartyRequest) (*genproto.PartyResponse, error) {
	query := `
		UPDATE party SET name = $2, slogan = $3, opened_at = $4, description = $5 WHERE id = $1
	`
	_, err := pt.db.Exec(query, id, req.Name, req.Slogan, req.OpenedAt, req.Description)
	if err != nil {
		return nil, err
	}

	return pt.SelectParty(&genproto.GetPartyInfoRequest{Id: id})
}

func (pt *PartyRepo) DeleteParty(id *genproto.DeletePartyRequest) (*genproto.Voidd, error) {
	query := `
		DELETE FROM party WHERE id = $1
	`
	_, err := pt.db.Exec(query, id.Id)
	if err != nil {
		return nil, err
	}

	return &genproto.Voidd{}, nil
}
