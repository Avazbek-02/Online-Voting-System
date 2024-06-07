package handler

import pb "api/genproto"

type Handler struct {
	ServiceParty pb.PartyServiceClient
	ServicePublic pb.PublicServiceClient
}

func NewHandler(party pb.PartyServiceClient, public pb.PublicServiceClient) *Handler {
	return &Handler{
		ServiceParty: party,
		ServicePublic: public,
	}
}