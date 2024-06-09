package handler

import (
	pb "api/genproto"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateVote(ctx *gin.Context)  {
	vote := pb.CreatePublicVoteRequest{}
	err := ctx.BindJSON(&vote)
	if err != nil{
		log.Fatal("error while vote bindjson !!",err)
		return
	}
	publicVoteRes, err := h.ServiceVote.CreatePublicVote(ctx,&vote)
	if err != nil{
		log.Fatal("Error while create publicVote in handler",err)
	}
	ctx.JSON(200,publicVoteRes)
}

func (h *Handler) UpdatePublicVote(ctx *gin.Context)  {
	id := ctx.Param("id")
	publicVote := pb.UpdatePublicVoteRequest{
		Id: id,
	}
	err := ctx.BindJSON(&publicVote)
	if err != nil{
		log.Fatal("Error while update publicVote in handler",err)
	}
	publicVoteRes, err := h.ServiceVote.UpdatePublicVote(ctx,&publicVote)
	if err != nil{
		log.Fatal("Error while update publicVote in handler",err)
	}
	ctx.JSON(200,publicVoteRes)
}

func (h *Handler) GetPublicVote(ctx *gin.Context)  {
	id := ctx.Param("id")
	publicVote := pb.GetPublicVoteInfoRequest{
		Id: id,
	}
	publicVoteRes,err := h.ServiceVote.GetPublicVoteInfo(ctx,&publicVote)
	if err != nil{
		log.Fatal("Error while get publicVote in handler",err)
	}
	ctx.JSON(200,publicVoteRes)
}

func (h *Handler) DeletePublicVote(ctx *gin.Context)  {
	id := ctx.Param("id")
	publicVote := pb.DeletePublicVoteRequest{
		Id: id,
	}
	_, err := h.ServiceVote.DeletePublicVote(ctx,&publicVote)
	if err != nil{
		log.Fatal("Error while delete publicVote in handler",err)
	}
	ctx.JSON(200,gin.H{"Deleted":"Successfully deleted"})
}