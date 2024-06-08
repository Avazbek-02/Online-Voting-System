package storage

import (
	v "public/genproto/genproto/public"

)

type StorageI interface {
	Party() Parties
	Public() Publics
}

type Parties interface {
	CreateParty(req *v.CreatePartyRequest) (*v.PartyResponse, error)
	SelectParty(req *v.GetPartyInfoRequest) (*v.PartyResponse, error)
	UpdateParty(req *v.UpdatePartyRequest) (*v.PartyResponse, error)
	DeleteParty(req *v.DeletePartyRequest) (*v.Voidd, error)
}

type Publics interface {
	CreatePublic(req *v.CreatePublicRequest) (*v.PublicResponse, error)
	SelectPublic(req *v.GetPublicInfoRequest) (*v.PublicResponse, error)
	UpdatePublic(req *v.UpdatePublicRequest) (*v.PublicResponse, error)
	DeletePublic(req *v.DeletePublicRequest) (*v.Void, error)
}
