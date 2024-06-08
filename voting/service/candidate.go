package service

import (
	v "all/genproto/genproto/voting"
	"all/storage"
	"context"
)

type CandidateService struct {
	v.UnimplementedCandidateServiceServer
	stg storage.StorageI
}

func NewCandidateService(stg storage.StorageI) *CandidateService {
	return &CandidateService{
		UnimplementedCandidateServiceServer: v.UnimplementedCandidateServiceServer{},
		stg:                                 stg,
	}
}

func (cs *CandidateService) CreateCandidate(ctx context.Context, req *v.CreateCandidateRequest) (*v.CandidateResponse, error) {
	candidate, err := cs.stg.Candidate().CreateCandidate(req)
	if err != nil {
		return nil, err
	}
	return candidate, nil
}

func (cs *CandidateService) GetCandidateInfo(ctx context.Context, req *v.GetCandidateInfoRequest) (*v.CandidateResponse, error) {
	candidate, err := cs.stg.Candidate().GetCandidateInfo(req)
	if err != nil {
		return nil, err
	}
	return candidate, nil
}

func (cs *CandidateService) UpdateCandidate(ctx context.Context, req *v.UpdateCandidateRequest) (*v.CandidateResponse, error) {
	candidate, err := cs.stg.Candidate().UpdateCandidate(req)
	if err != nil {
		return nil, err
	}
	return candidate, nil
}

func (cs *CandidateService) DeleteCandidate(ctx context.Context, req *v.DeleteCandidateRequest) (*v.Void1, error) {
	_, err := cs.stg.Candidate().DeleteCandidate(req)
	if err != nil {
		return nil, err
	}
	return &v.Void1{}, nil
}
