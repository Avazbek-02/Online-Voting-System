package service

import (
	"context"
	// "log"
	pb "public_service/genproto"
	"public_service/storage/postgres"
)

type partyService struct{
	stg *postgres.Storage
	pb.UnimplementedPartyServiceServer
}

func NewPartyService(stg *postgres.Storage) *partyService {
	return &partyService{stg: stg}
}

func (pt *partyService) CreateBasket(ctx context.Context,req *pb.CreatePartyRequest) (*pb.GetPartyInfoRequest,error) {
	// var party pb.GetPartyInfoRequest
	// party,err := pt.stg.Party.CreateParty(req)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	return nil, nil
}