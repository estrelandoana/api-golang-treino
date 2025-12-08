package main

import (
	"log"
	"net/http"

	"github.com/estrelandoana/api-golang-treino/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter() //cria um roteador

	r.HandleFunc("/musicas", handler.ListarMusicas).Methods("GET")
	r.HandleFunc("/musicas", handler.CreateMusica).Methods("POST") //def e associa rotas x handlers
	r.HandleFunc("/musicas/{id}", handler.GetMusica).Methods("GET")

	log.Println("Servidor porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
