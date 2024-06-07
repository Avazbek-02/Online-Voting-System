package storage

import (
	v "all/voting/genproto/genproto/voting"
)

type StorageI interface {
	Candidate() Candidates
	Election() Elections
	PublicVote() PublicVotes
}

type Candidates interface {
	CreateCandidate(req *v.CreateCandidateRequest) (*v.CandidateResponse, error)
	SelectCandidate(id *v.GetCandidateInfoRequest) (*v.CandidateResponse, error)
	UpdateCandidate(req *v.UpdateCandidateRequest) (*v.CandidateResponse, error)
	DeleteCandidate(id *v.DeleteCandidateRequest) (*v.Void1, error)
}

type Elections interface {
	CreateElection(req *v.CreateElectionRequest) (*v.ElectionResponse, error)
	SelectElection(id *v.GetElectionInfoRequest) (*v.ElectionResponse, error)
	UpdateElection(req *v.UpdateElectionRequest) (*v.ElectionResponse, error)
	DeleteElection(id *v.DeleteElectionRequest) (*v.Void2, error)
}

type PublicVotes interface {
	CreatePublicVote(req *v.CreatePublicVoteRequest) (*v.PublicVoteResponse, error)
	SelectPublicVote(id *v.GetPublicVoteInfoRequest) (*v.PublicVoteResponse, error)
	UpdatePublicVote(req *v.UpdatePublicVoteRequest) (*v.PublicVoteResponse, error)
	DeletePublicVote(id *v.DeletePublicVoteRequest) (*v.Void3, error)
}
