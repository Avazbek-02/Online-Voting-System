package service

import (
	v "all/voting/genproto/genproto/voting"
	"all/voting/storage"
)

type CandidateService struct {
	stg storage.StorageI
}

func NewCandidateService(stg storage.StorageI) *CandidateService {
	return &CandidateService{stg: stg}
}

func (cs *CandidateService) CreateCandidate(req *v.CreateCandidateRequest) (*v.CandidateResponse, error) {
	candidate, err := cs.stg.Candidate().CreateCandidate(req)
	if err != nil {
		return nil, err
	}
	return candidate, err
}

func (cs *CandidateService) GetCandidateInfo(req *v.GetCandidateInfoRequest) (*v.CandidateResponse, error) {
	candidate, err := cs.stg.Candidate().SelectCandidate(req)
	if err != nil {
		return nil, err
	}
	return candidate, err
}

func (cs *CandidateService) UpdateCandidate(req *v.UpdateCandidateRequest) (*v.CandidateResponse, error) {
	candidate, err := cs.stg.Candidate().UpdateCandidate(req)
	if err != nil {
		return nil, err
	}
	return candidate, err
}

func (cs *CandidateService) DeleteCandidate(req *v.DeleteCandidateRequest) (*v.Void1, error) {
	_, err := cs.stg.Candidate().DeleteCandidate(req)
	if err != nil {
		return nil, err
	}
	return &v.Void1{}, err
}
