package postgres

import (
	"database/sql"
	"log"

	pb "public_service/genproto"

	"github.com/satori/go.uuid"
)

type PartyRepo struct {
	db *sql.DB
}

func NewParyRepo(db *sql.DB) *PartyRepo {
	return &PartyRepo{db: db}
}

func (pt *PartyRepo) CreateParty(party *pb.CreatePartyRequest) (*pb.CreatePartyRequest,error) {
	query := `
		insert into party(id,name,slogan,opened_date,description) values($1,$2,$3,$4)
	`
	id := uuid.NewV4()

	_, err := pt.db.Exec(query,id, party.Name, party.Slogan, party.OpenedAt, party.Description)
	if err != nil {
		log.Fatal("Error while insert paryt")
	}
	query = `
		select id,name,slogan,opened,description from paryt where id = $1
	`
	var backparty pb.GetPartyInfoRequest
	err = pt.db.QueryRow()
	return party,err
}
