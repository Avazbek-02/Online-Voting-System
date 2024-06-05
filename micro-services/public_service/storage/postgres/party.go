package postgres

import (
	pb "all/micro-services/public_service/genproto/Public"
	"database/sql"
)

type PartyRepo struct {
	db *sql.DB
}

func NewParyRepo(db *sql.DB) *PartyRepo {
	return &PartyRepo{db: db}
}

func (pt *PartyRepo) CreateParty(party *pb.CreatePartyRequest) {

}
