package handler

import (
	pb "api/genproto"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPublicById(ctx *gin.Context)  {
	id := ctx.Param("id")
	public := pb.GetPublicInfoRequest{
		Id: id,
	}
	publicRes, err := h.ServicePublic.GetPublicInfo(ctx,&public)
	if err != nil{
		log.Fatal("error while get public in handler",err)
		return
	}
	ctx.JSON(200,publicRes)
}

func (h *Handler) CreatePublic(ctx *gin.Context)  {
	public := pb.CreatePublicRequest{}
	err := ctx.BindJSON(&public)
	if err != nil{
		log.Fatal("Error while public create bindjson !",err)
		return
	}
	publicRes,err := h.ServicePublic.CreatePublic(ctx,&public)
	if err != nil{
		log.Fatal("error while create public in handler",err)
		return
	}
	ctx.JSON(200,publicRes)
}

func (h *Handler) DeletePublic(ctx *gin.Context)  {
	id := ctx.Param("id")
	public := pb.DeletePublicRequest{}
	public.Id = id

	_,err := h.ServicePublic.DeletePublic(ctx,&public)
	if err != nil{
		log.Fatal("Error while delete public in handler")
		return
	}
	ctx.JSON(200,gin.H{"Deleted":"Successfully deleted"})
}

func (h *Handler) UpdatePublic(ctx *gin.Context)  {
	id := ctx.Param("id")
	public := pb.UpdatePublicRequest{}
	public.Id = id
	err := ctx.BindJSON(&public)
	if err != nil{
		log.Fatal("error while update public in handler")
		return
	}
	publicRes, err := h.ServicePublic.UpdatePublic(ctx,&public)
	if err != nil{
		log.Fatal("Error while update public in handler")
		return
	}
	ctx.JSON(200,publicRes)
}