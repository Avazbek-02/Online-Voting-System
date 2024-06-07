package service

import (
	"all/public/genproto/genproto"
	"all/public/storage"
	"context"
)

type PartyService struct {
	stg storage.StorageI
	genproto.UnimplementedPartyServiceServer
}

func NewPartyService(stg storage.StorageI) *PartyService {
	return &PartyService{stg: stg}
}

func (pt *PartyService) CreateParty(ctx context.Context, req *genproto.CreatePartyRequest) (*genproto.PartyResponse, error) {
	party, err := pt.stg.Party().CreateParty(req)
	if err != nil {
		return nil, err
	}
	return party, nil
}

func (pt *PartyService) GetParty(ctx context.Context, req *genproto.GetPartyInfoRequest) (*genproto.PartyResponse, error) {
	party, err := pt.stg.Party().SelectParty(req)
	if err != nil {
		return nil, err
	}
	return party, nil
}

func (pt *PartyService) UpdateParty(ctx context.Context, req *genproto.UpdatePartyRequest) (*genproto.PartyResponse, error) {
	party, err := pt.stg.Party().UpdateParty(req)
	if err != nil {
		return nil, err
	}
	return party, nil
}

func (pt *PartyService) DeleteParty(ctx context.Context, req *genproto.DeletePartyRequest) (*genproto.Voidd, error) {
	_, err := pt.stg.Party().DeleteParty(req)
	if err != nil {
		return nil, err
	}
	return &genproto.Voidd{}, nil
}
