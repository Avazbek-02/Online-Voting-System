package handler

import (
	pb "api/genproto"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCondidateById(ctx *gin.Context) {

	
}

func (h *Handler) CreateCondidate(ctx *gin.Context) {
	condidate := pb.CandidateResponse{}
	err := ctx.BindJSON(&condidate)
	if err != nil{
		log.Fatal("error while condidate bindjson",err)
		return
	}
}

func (h *Handler) DeleteCondidate(ctx *gin.Context) {

}

func (h *Handler) UpdateCondidate(ctx *gin.Context) {

}
