package postgres

import (
	"all/public/genproto/genproto"
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

type PublicRepo struct {
	db *sql.DB
}

func NewPublicRepo(db *sql.DB) *PublicRepo {
	return &PublicRepo{db: db}
}

func (pr *PublicRepo) CreatePublic(req *genproto.CreatePublicRequest) (*genproto.PublicResponse, error) {
	query := `
		INSERT INTO public(id, firstname, lastname, birthday, gender, nationality, party_id) 
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`
	id := uuid.NewV4().String()

	_, err := pr.db.Exec(query, id, req.Firstname, req.Lastname, req.Birthday, req.Gender, req.Nationality, req.PartyId)
	if err != nil {
		return nil, err
	}

	return &genproto.PublicResponse{
		Id:          id,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Birthday:    req.Birthday,
		Gender:      req.Gender,
		Nationality: req.Nationality,
		PartyId:     req.PartyId,
	}, nil
}

func (pr *PublicRepo) SelectPublic(id *genproto.GetPublicInfoRequest) (*genproto.PublicResponse, error) {
	query := `
		SELECT id, firstname, lastname, birthday, gender, nationality, party_id FROM public WHERE id = $1
	`
	var public genproto.PublicResponse
	err := pr.db.QueryRow(query, id.Id).Scan(
		&public.Id,
		&public.Firstname,
		&public.Lastname,
		&public.Birthday,
		&public.Gender,
		&public.Nationality,
		&public.PartyId,
	)
	if err != nil {
		return nil, err
	}

	return &public, nil
}

func (pr *PublicRepo) UpdatePublic(id string, req *genproto.CreatePublicRequest) (*genproto.PublicResponse, error) {
	query := `
		UPDATE public SET firstname = $2, lastname = $3, birthday = $4, gender = $5, nationality = $6, party_id = $7 WHERE id = $1
	`
	_, err := pr.db.Exec(query, id, req.Firstname, req.Lastname, req.Birthday, req.Gender, req.Nationality, req.PartyId)
	if err != nil {
		return nil, err
	}

	return pr.SelectPublic(&genproto.GetPublicInfoRequest{Id: id})
}

func (pr *PublicRepo) DeletePublic(id string) error {
	query := `
		DELETE FROM public WHERE id = $1
	`
	_, err := pr.db.Exec(query, id)
	return err
}
