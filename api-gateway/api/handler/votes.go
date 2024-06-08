package handler

import (
	pb "api/genproto"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateVote(ctx *gin.Context)  {
	vote := pb.PublicVoteResponse{}
	err := ctx.BindJSON(&vote)
	if err != nil{
		log.Fatal("error while vote bindjson !!",err)
		return
	}
}