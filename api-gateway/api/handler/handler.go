package handler

import pb "api/genproto"

type Handler struct {
	ServiceParty pb.PartyServiceClient
	ServicePublic pb.PublicServiceClient
	ServiceCandidate pb.CandidateServiceClient
	ServiceElection pb.ElectionServiceClient
	ServiceVote pb.PublicVoteServiceClient
}

func NewHandler(party pb.PartyServiceClient, public pb.PublicServiceClient) *Handler {
	return &Handler{
		ServiceParty: party,
		ServicePublic: public,
	}
}