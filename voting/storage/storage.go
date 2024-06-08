package storage

import "all/voting/genproto/genproto/voting"

type StorageI interface {
	Candidate() Candidates
	Election() Elections
	PublicVote() PublicVotes
}

type Candidates interface {
	CreateCandidate(req *voting.CreateCandidateRequest) (*voting.CandidateResponse, error)
	GetCandidateInfo(req *voting.GetCandidateInfoRequest) (*voting.CandidateResponse, error)
	UpdateCandidate(req *voting.UpdateCandidateRequest) (*voting.CandidateResponse, error)
	DeleteCandidate(req *voting.DeleteCandidateRequest) (*voting.Void1, error)
}

type Elections interface {
	CreateElection(req *voting.CreateElectionRequest) (*voting.ElectionResponse, error)
	GetElectionInfo(req *voting.GetElectionInfoRequest) (*voting.ElectionResponse, error)
	UpdateElection(req *voting.UpdateElectionRequest) (*voting.ElectionResponse, error)
	DeleteElection(req *voting.DeleteElectionRequest) (*voting.Void2, error)
}

type PublicVotes interface {
	CreatePublicVote(req *voting.CreatePublicVoteRequest) (*voting.PublicVoteResponse, error)
	GetPublicVoteInfo(req *voting.GetPublicVoteInfoRequest) (*voting.PublicVoteResponse, error)
	UpdatePublicVote(req *voting.UpdatePublicVoteRequest) (*voting.PublicVoteResponse, error)
	DeletePublicVote(req *voting.DeletePublicVoteRequest) (*voting.Void3, error)
}
