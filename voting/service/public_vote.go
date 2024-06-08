package service

import (
	v "voting/genproto/genproto/voting"
	"voting/storage"
	"context"
)

type PublicVoteService struct {
	v.UnimplementedPublicVoteServiceServer
	stg storage.StorageI
}

func NewPublicVoteService(stg storage.StorageI) *PublicVoteService {
	return &PublicVoteService{
		UnimplementedPublicVoteServiceServer: v.UnimplementedPublicVoteServiceServer{},
		stg:                                  stg,
	}
}

func (ps *PublicVoteService) CreatePublicVote(ctx context.Context, req *v.CreatePublicVoteRequest) (*v.PublicVoteResponse, error) {
	publicVote, err := ps.stg.PublicVote().CreatePublicVote(req)
	if err != nil {
		return nil, err
	}
	return publicVote, nil
}

func (ps *PublicVoteService) GetPublicVoteInfo(ctx context.Context, req *v.GetPublicVoteInfoRequest) (*v.PublicVoteResponse, error) {
	publicVote, err := ps.stg.PublicVote().GetPublicVoteInfo(req)
	if err != nil {
		return nil, err
	}
	return publicVote, nil
}

func (ps *PublicVoteService) UpdatePublicVote(ctx context.Context, req *v.UpdatePublicVoteRequest) (*v.PublicVoteResponse, error) {
	publicVote, err := ps.stg.PublicVote().UpdatePublicVote(req)
	if err != nil {
		return nil, err
	}
	return publicVote, nil
}

func (ps *PublicVoteService) DeletePublicVote(ctx context.Context, req *v.DeletePublicVoteRequest) (*v.Void3, error) {
	_, err := ps.stg.PublicVote().DeletePublicVote(req)
	if err != nil {
		return nil, err
	}
	return &v.Void3{}, nil
}
