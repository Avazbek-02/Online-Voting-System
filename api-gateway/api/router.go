package api

import (
	"api/api/handler"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewEngine(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/swager/*any", ginSwagger.WrapHandler(files.Handler))
	p := r.Group("party")
	p.POST("create", h.CreateParty)
	return r
}