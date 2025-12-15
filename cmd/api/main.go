package main

import (
	_ "github.com/estrelandoana/api-golang-treino/docs"
	"github.com/estrelandoana/api-golang-treino/internal/db"
	"github.com/estrelandoana/api-golang-treino/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API de Músicas (Ana)
// @version         1.0
// @description     API REST para gerenciamento de músicas.
// @host            localhost:8080
// @BasePath        /api/v1
// @schemes         http
func main() {
	db.ConectorDB()
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/musicas", handler.GinListarMusicas)
		api.POST("/musicas", handler.GinCreateMusica)
		api.GET("/musicas/:id", handler.GinGetMusica)
		api.PUT("/musicas/:id", handler.GinUpdateMusica)
		api.DELETE("/musicas/:id", handler.GinDeleteMusica)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
