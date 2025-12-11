package main

import (
	"log"
	"net/http"

	"github.com/estrelandoana/api-golang-treino/internal/db"
	"github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/estrelandoana/api-golang-treino/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	db.ConectorDB()
	db.DB.AutoMigrate(&entity.Musica{})
	r := mux.NewRouter() //cria um roteador

	r.HandleFunc("/musicas", handler.ListarMusicas).Methods("GET")
	r.HandleFunc("/musicas", handler.CreateMusica).Methods("POST") //def e associa rotas x handlers
	r.HandleFunc("/musicas/{id}", handler.GetMusica).Methods("GET")
	r.HandleFunc("/musicas/{id}", handler.UpdateMusica).Methods("PUT")
	r.HandleFunc("/musicas/{id}", handler.DeleteMusica).Methods("DELETE")

	log.Println("Servidor porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
