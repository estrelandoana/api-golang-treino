package router

import (
	"github.com/estrelandoana/api-golang-treino/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/musicas", controller.ListarMusica)
		api.POST("/musicas", controller.CreateMusica)
		api.GET("/musicas/:id", controller.GetMusica)
		api.PUT("/musicas/:id", controller.UpdateMusica)
		api.DELETE("/musicas/:id", controller.DeleteMusica)
	}
	return r
}
