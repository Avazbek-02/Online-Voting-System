package handler

import (
	pb "api/genproto"
	"log"

	"github.com/gin-gonic/gin"//
)

func (h *Handler) GetCondidateById(ctx *gin.Context) {
	id := ctx.Param("id")
	condidate := pb.GetCandidateInfoRequest{
		Id: id,
	}
	condidateRes, err := h.ServiceCandidate.GetCandidateInfo(ctx,&condidate)
	if err != nil{
		log.Fatal("Error while get condidate in handler",err)
	}
	ctx.JSON(200,condidateRes)
}

func (h *Handler) CreateCondidate(ctx *gin.Context) {
	condidate := pb.CreateCandidateRequest{}
	err := ctx.BindJSON(&condidate)
	if err != nil{
		log.Fatal("error while create condidate bindjson in handler",err)
		return
	}
	condidateRes, err := h.ServiceCandidate.CreateCandidate(ctx,&condidate)
	if err != nil{
		log.Fatal("Error while create condidate in handler",err)
	}
	ctx.JSON(200,condidateRes)
}

func (h *Handler) DeleteCondidate(ctx *gin.Context) {
	id := ctx.Param("id")
	condidate := pb.DeleteCandidateRequest{
		Id: id,
	}
	_,err := h.ServiceCandidate.DeleteCandidate(ctx,&condidate)
	if err != nil{
		log.Fatal("Error while delete condidate in handler",err)
	}
	ctx.JSON(200,gin.H{"Deleted":"Succsessfully deleted"})
}

func (h *Handler) UpdateCondidate(ctx *gin.Context) {
	id := ctx.Param("id")
	condidate := pb.UpdateCandidateRequest{
		Id: id,
	}
	err := ctx.BindJSON(&condidate)
	if err != nil{
		log.Fatal("Error while update condidate in handler",err)
	}
	condidateRes, err := h.ServiceCandidate.UpdateCandidate(ctx,&condidate)
	if err != nil{
		log.Fatal("Error while update condidate in handler",err)
	}
	ctx.JSON(200,condidateRes)
}
