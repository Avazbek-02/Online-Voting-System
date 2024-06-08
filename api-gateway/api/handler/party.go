package handler

import (
	pb "api/genproto"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateParty(ctx *gin.Context)  {
	party := pb.CreatePartyRequest{}
	
	fmt.Println("Worked")
	err := ctx.BindJSON(&party)
	
	if err != nil{
		log.Fatal("Error while create party read json",err)
		return
	}
	h.ServiceParty.CreateParty(ctx, &party)
	ctx.JSON(200,&party)
}

func (h *Handler) UpdateParty(ctx *gin.Context)  {
	// id := ctx.Param("id")
	party := pb.PartyResponse{}
	err := ctx.BindJSON(&party)
	if err != nil{
		log.Fatal("Error while update party and read body",err)
		return
	}
	ctx.JSON(200,&party)
}

func (h *Handler) DeleteParty(ctx *gin.Context)  {
	// id := ctx.Param("id")
	ctx.JSON(200,gin.H{"Deleted":"Party is deleted"})
}

func (h *Handler) GetPartyById(ctx *gin.Context)  {
	// id := ctx.Param("id")

}