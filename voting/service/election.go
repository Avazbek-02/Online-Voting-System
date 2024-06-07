package service

import (
	v "all/voting/genproto/genproto/voting"
	"all/voting/storage"
)

type ElectionService struct {
	storage storage.StorageI
}

func NewElectionService(storage storage.StorageI) *ElectionService {
	return &ElectionService{storage: storage}
}

func (es *ElectionService) CreateElection(req *v.CreateElectionRequest) (*v.ElectionResponse, error) {
	election, err := es.storage.Election().CreateElection(req)
	if err != nil {
		return nil, err
	}
	return election, nil
}

func (es *ElectionService) GetElectionInfo(req *v.GetElectionInfoRequest) (*v.ElectionResponse, error) {
	election, err := es.storage.Election().SelectElection(req)
	if err != nil {
		return nil, err
	}
	return election, nil
}

func (es *ElectionService) UpdateElection(req *v.UpdateElectionRequest) (*v.ElectionResponse, error) {
	election, err := es.storage.Election().UpdateElection(req)
	if err != nil {
		return nil, err
	}
	return election, nil
}

func (es *ElectionService) DeleteElection(req *v.DeleteElectionRequest) (*v.Void2, error) {
	_, err := es.storage.Election().DeleteElection(req)
	if err != nil {
		return nil, err
	}
	return &v.Void2{}, nil
}
