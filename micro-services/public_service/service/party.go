package service

import (
	pb "all/micro-services/public_service/genproto/Public"
	"all/micro-services/public_service/storage/postgres"
	"context"
)

type partyService struct{
	stg *postgres.Storage
	pb.UnimplementedPartyServiceServer
}

func NewPartyService(stg *postgres.Storage) *partyService {
	return &partyService{stg: stg}
}

func (pt *partyService) CreateBasket(ctx context.Context,req *pb.CreatePartyRequest) (*pb.GetPartyInfoRequest,error) {
	res,err := pt.stg.Party
}