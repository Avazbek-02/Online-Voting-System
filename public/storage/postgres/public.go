package postgres

import (
	v "public/genproto/genproto/public"
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

type PublicRepo struct {
	db *sql.DB
}

func NewPublicRepo(db *sql.DB) *PublicRepo {
	return &PublicRepo{db: db}
}

func (pr *PublicRepo) CreatePublic(req *v.CreatePublicRequest) (*v.PublicResponse, error) {
	query := `
		INSERT INTO public(id, firstname, lastname, birthday, gender, nationality, party_id) 
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`
	id := uuid.NewV4().String()

	_, err := pr.db.Exec(query, id, req.Firstname, req.Lastname, req.Birthday, req.Gender, req.Nationality, req.PartyId)
	if err != nil {
		return nil, err
	}

	return &v.PublicResponse{
		Id:          id,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Birthday:    req.Birthday,
		Gender:      req.Gender,
		Nationality: req.Nationality,
		PartyId:     req.PartyId,
	}, nil
}

func (pr *PublicRepo) SelectPublic(id *v.GetPublicInfoRequest) (*v.PublicResponse, error) {
	query := `
		SELECT id, firstname, lastname, birthday, gender, nationality, party_id FROM public WHERE id = $1
	`
	var public v.PublicResponse
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

func (pr *PublicRepo) UpdatePublic(req *v.UpdatePublicRequest) (*v.PublicResponse, error) {
	query := `
		UPDATE public SET firstname = $2, lastname = $3, birthday = $4, gender = $5, nationality = $6, party_id = $7 WHERE id = $1
	`
	_, err := pr.db.Exec(query, req.Id, req.Firstname, req.Lastname, req.Birthday, req.Gender, req.Nationality, req.PartyId)
	if err != nil {
		return nil, err
	}

	return pr.SelectPublic(&v.GetPublicInfoRequest{Id: req.Id})
}

func (pr *PublicRepo) DeletePublic(id *v.DeletePublicRequest) (*v.Void, error) {
	query := `
		DELETE FROM public WHERE id = $1
	`
	_, err := pr.db.Exec(query, id.Id)
	if err != nil {
		return nil, err
	}
	return &v.Void{}, nil
}
