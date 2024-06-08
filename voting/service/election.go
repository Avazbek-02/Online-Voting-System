package service

import (
	v "all/genproto/genproto/voting"
	"all/storage"
	"context"
)

type ElectionService struct {
	v.UnimplementedElectionServiceServer
	stg storage.StorageI
}

func NewElectionService(stg storage.StorageI) *ElectionService {
	return &ElectionService{
		UnimplementedElectionServiceServer: v.UnimplementedElectionServiceServer{},
		stg:                                 stg,
	}
}

func (es *ElectionService) CreateElection(ctx context.Context, req *v.CreateElectionRequest) (*v.ElectionResponse, error) {
	election, err := es.stg.Election().CreateElection(req)
	if err != nil {
		return nil, err
	}
	return election, nil
}

func (es *ElectionService) GetElectionInfo(ctx context.Context, req *v.GetElectionInfoRequest) (*v.ElectionResponse, error) {
	election, err := es.stg.Election().GetElectionInfo(req)
	if err != nil {
		return nil, err
	}
	return election, nil
}

func (es *ElectionService) UpdateElection(ctx context.Context, req *v.UpdateElectionRequest) (*v.ElectionResponse, error) {
	election, err := es.stg.Election().UpdateElection(req)
	if err != nil {
		return nil, err
	}
	return election, nil
}

func (es *ElectionService) DeleteElection(ctx context.Context, req *v.DeleteElectionRequest) (*v.Void2, error) {
	_, err := es.stg.Election().DeleteElection(req)
	if err != nil {
		return nil, err
	}
	return &v.Void2{}, nil
}
