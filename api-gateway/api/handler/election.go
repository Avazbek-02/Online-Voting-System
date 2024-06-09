package handler

import (
	pb "api/genproto"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateElection(ctx *gin.Context) {
	election := pb.CreateElectionRequest{}
	err := ctx.BindJSON(&election)
	if err != nil {
		log.Fatal("error while election check it !!", err)
		return
	}
	electionRes, err := h.ServiceElection.CreateElection(ctx, &election)
	if err != nil {
		log.Fatal("Error while create election in handler", err)
		return
	}
	ctx.JSON(200, electionRes)
}

func (h *Handler) UpdateElection(ctx *gin.Context) {
	id := ctx.Param("id")
	election := pb.UpdateElectionRequest{
		Id: id,
	}
	err := ctx.BindJSON(&election)
	if err != nil {
		log.Fatal("Error while updaate election in handler", err)
		return
	}
	electionRes, err := h.ServiceElection.UpdateElection(ctx, &election)
	if err != nil {
		log.Fatal("Error while update election in handler", err)
		return
	}
	ctx.JSON(200, electionRes)
}

func (h *Handler) GetElectoin(ctx *gin.Context) {
	id := ctx.Param("id")
	election := pb.GetElectionInfoRequest{
		Id: id,
	}
	err := ctx.BindJSON(&election)
	if err != nil {
		log.Fatal("Error while get election in handler", err)
	}
	electionRes, err := h.ServiceElection.GetElectionInfo(ctx, &election)
	if err != nil {
		log.Fatal("Error while get election in handler", err)
	}
	ctx.JSON(200, electionRes)
}

func (h *Handler) DeleteElection(ctx *gin.Context) {
	id := ctx.Param("id")
	election := pb.DeleteElectionRequest{
		Id: id,
	}
	_, err := h.ServiceElection.DeleteElection(ctx,&election)
	if err != nil{
		log.Fatal("Error while delete election in handler",err)
	}
	ctx.JSON(200,gin.H{"Deleted":"Succsessfully deleted"})
}
