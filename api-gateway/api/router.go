package api

import (
	"api/api/handler"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewEngine(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/swager/*any", ginSwagger.WrapHandler(files.Handler))

	party := router.Group("party")
	party.GET("get/:id",h.GetPartyById)
	party.POST("create", h.CreateParty)
	party.DELETE("delete/:id",h.DeleteParty)
	party.PUT("update/:id",h.UpdateParty)

	public := router.Group("public")
	public.GET("get/:id",h.GetPublicById)
	public.POST("create",h.CreatePublic)
	public.PUT("update/:id",h.UpdatePublic)
	public.DELETE("delete/:id",h.DeletePublic)

	candidate := router.Group("candidate")
	candidate.GET("get/:id",h.GetCondidateById)
	candidate.POST("create",h.GetCondidateById)
	candidate.PUT("update/:id",h.GetCondidateById)
	candidate.DELETE("update/:id",h.DeleteCondidate)

	election := router.Group("election")
	election.POST("create",h.CreateElection)
	election.GET("get/:id",h.GetElectoin)
	election.PUT("update",h.UpdateElection)
	election.DELETE("delete",h.DeleteElection)

	return router
}
