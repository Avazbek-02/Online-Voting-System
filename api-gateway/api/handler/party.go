package handler

import (
	pb "api/genproto"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateParty(ctx *gin.Context) {
	party := pb.CreatePartyRequest{}

	err := ctx.BindJSON(&party)
	if err != nil {
		log.Fatal("Error while create party read json", err)
		return
	}
	partyRes, err := h.ServiceParty.CreateParty(ctx, &party)
	if err != nil {
		log.Fatal("Error while create party in handler", err)
		return
	}
	ctx.JSON(200, partyRes)
}

func (h *Handler) UpdateParty(ctx *gin.Context) {
	id := ctx.Param("id")
	party := pb.UpdatePartyRequest{}
	party.Id = id
	err := ctx.BindJSON(&party)
	if err != nil {
		log.Fatal("Error while update party and read body", err)
		return
	}
	backparty, err := h.ServiceParty.UpdateParty(ctx, &party)
	if err != nil {
		log.Fatal("error while update in handler", err)
		return
	}
	ctx.JSON(200, backparty)
}

func (h *Handler) DeleteParty(ctx *gin.Context) {
	id := ctx.Param("id")
	party := pb.DeletePartyRequest{
		Id: id,
	}
	_, err := h.ServiceParty.DeleteParty(ctx, &party)
	if err != nil {
		log.Fatal("Error while delete in handler", err)
		return
	}
	ctx.JSON(200, gin.H{"Deleted": "Party is deleted"})
}

func (h *Handler) GetPartyById(ctx *gin.Context) {
	id := ctx.Param("id")
	party := pb.GetPartyInfoRequest{
		Id: id,
	}
	getInfo, err := h.ServiceParty.GetPartyInfo(ctx, &party)
	if err != nil {
		log.Fatal("error while get info in handler", err)
		return
	}
	ctx.JSON(200, getInfo)
}
