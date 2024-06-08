package service

import (
	genproto "all/public/genproto/genproto/public"
	"all/public/storage"
	"context"
)

type PublicService struct {
	stg storage.StorageI
	genproto.UnimplementedPublicServiceServer
}

func NewPublicService(stg storage.StorageI) *PublicService {
	return &PublicService{stg: stg}
}

func (ps *PublicService) CreatePublic(ctx context.Context, req *genproto.CreatePublicRequest) (*genproto.PublicResponse, error) {
	public, err := ps.stg.Public().CreatePublic(req)
	if err != nil {
		return nil, err
	}
	return public, nil
}

func (ps *PublicService) GetPublic(ctx context.Context, req *genproto.GetPublicInfoRequest) (*genproto.PublicResponse, error) {
	public, err := ps.stg.Public().SelectPublic(req)
	if err != nil {
		return nil, err
	}
	return public, nil
}

func (ps *PublicService) UpdatePublic(ctx context.Context, req *genproto.UpdatePublicRequest) (*genproto.PublicResponse, error) {
	public, err := ps.stg.Public().UpdatePublic(req)
	if err != nil {
		return nil, err
	}
	return public, nil
}

func (ps *PublicService) DeletePublic(ctx context.Context, req *genproto.DeletePublicRequest) (*genproto.Void, error) {
	_, err := ps.stg.Public().DeletePublic(req)
	if err != nil {
		return nil, err
	}
	return &genproto.Void{}, nil
}
