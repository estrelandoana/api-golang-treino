package main

import (
	"fmt"

	_ "github.com/estrelandoana/api-golang-treino/docs"
	"github.com/estrelandoana/api-golang-treino/internal/database"
	"github.com/estrelandoana/api-golang-treino/internal/router"
)

// @title           API de Músicas (Ana)
// @version         1.0
// @description     API REST para gerenciamento de músicas.
// @host            localhost:8080
// @BasePath        /api/v1
// @schemes         http
func main() {
	database.ConectorDB()
	r := router.SetupRouter()
	port := ":8080"
	fmt.Printf("Rodando em http://localhost%s\n", port)
	r.Run(port)
}
