package storage

import (
	"all/public/genproto/genproto"
)

type StorageI interface {
	Party() Parties
	Public() Publics
}

type Parties interface {
	CreateParty(req *genproto.CreatePartyRequest) (*genproto.PartyResponse, error)
	SelectParty(id *genproto.GetPartyInfoRequest) (*genproto.PartyResponse, error)
	UpdateParty(req *genproto.UpdatePartyRequest) (*genproto.PartyResponse, error)
	DeleteParty(id *genproto.DeletePartyRequest) (*genproto.Voidd, error)
}

type Publics interface {
	CreatePublic(req *genproto.CreatePublicRequest) (*genproto.PublicResponse, error)
	SelectPublic(id *genproto.GetPublicInfoRequest) (*genproto.PublicResponse, error)
	UpdatePublic(req *genproto.UpdatePublicRequest) (*genproto.PublicResponse, error)
	DeletePublic(id *genproto.DeletePublicRequest) (*genproto.Void, error)
}
