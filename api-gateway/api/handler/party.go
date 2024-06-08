package handler

import (
	pb "api/genproto"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateParty(ctx *gin.Context)  {
	id := uuid.New()
	party := pb.PartyResponse{}
	party.Id = id.String()
	
	fmt.Println("Worked")
	err := ctx.BindJSON(&party)
	if err != nil{
		log.Fatal("Error while create party read json",err)
		return
	}
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