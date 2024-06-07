package service

import (
	v "all/voting/genproto/genproto/voting"
	"all/voting/storage"
)

type PublicVoteService struct {
	storage storage.StorageI
}

func NewPublicVoteService(storage storage.StorageI) *PublicVoteService {
	return &PublicVoteService{storage: storage}
}

func (ps *PublicVoteService) CreatePublicVote(req *v.CreatePublicVoteRequest) (*v.PublicVoteResponse, error) {
	publicVote, err := ps.storage.PublicVote().CreatePublicVote(req)
	if err != nil {
		return nil, err
	}
	return publicVote, nil
}

func (ps *PublicVoteService) GetPublicVoteInfo(req *v.GetPublicVoteInfoRequest) (*v.PublicVoteResponse, error) {
	publicVote, err := ps.storage.PublicVote().SelectPublicVote(req)
	if err != nil {
		return nil, err
	}
	return publicVote, nil
}

func (ps *PublicVoteService) UpdatePublicVote(req *v.UpdatePublicVoteRequest) (*v.PublicVoteResponse, error) {
	publicVote, err := ps.storage.PublicVote().UpdatePublicVote(req)
	if err != nil {
		return nil, err
	}
	return publicVote, nil
}

func (ps *PublicVoteService) DeletePublicVote(req *v.DeletePublicVoteRequest) (*v.Void3, error) {
	_, err := ps.storage.PublicVote().DeletePublicVote(req)
	if err != nil {
		return nil, err
	}
	return &v.Void3{}, nil
}
