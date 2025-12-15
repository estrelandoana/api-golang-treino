package main

import (
	"github.com/estrelandoana/api-golang-treino/internal/db"
	"github.com/estrelandoana/api-golang-treino/internal/handler"
	"github.com/gin-gonic/gin"
)

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
	r.Run(":8080")
}
