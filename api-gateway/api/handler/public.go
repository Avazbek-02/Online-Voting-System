package handler

import (
	pb "api/genproto"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPublicById(ctx *gin.Context)  {
	// id := ctx.Param("id")

}

func (h *Handler) CreatePublic(ctx *gin.Context)  {
	public := pb.PublicResponse{}
	err := ctx.BindJSON(&public)
	if err != nil{
		log.Fatal("Error while public create bindjson !",err)
		return
	}
	
}

func (h *Handler) DeletePublic(ctx *gin.Context)  {
	
}

func (h *Handler) UpdatePublic(ctx *gin.Context)  {
	
}