package handler

import (
	pb "api/genproto"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateElection(ctx *gin.Context) {
	election := pb.ElectionResponse{}
	err := ctx.BindJSON(&election)
	if err != nil{
		log.Fatal("error while election check it !!",err)
		return
	}
}

func (h *Handler) UpdateElection(ctx *gin.Context)  {
	
}

func (h *Handler) GetElectoin(ctx *gin.Context)  {
	
}

func (h *Handler) DeleteElection(ctx *gin.Context)  {
	
}