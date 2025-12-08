package handler

import (
	"encoding/json"
	"net/http"

	model "github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var musicas = []model.Musica{
	{
		ID:      "1",
		Titulo:  "Nome da musica",
		Artista: "Artista da musica",
		Album:   "Album da musica",
		Ano:     2022,
		Genero:  "Genero da musica",
		Duracao: 30,
	},
} //slice

func ListarMusicas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musicas)
} //listar

func CreateMusica(w http.ResponseWriter, r *http.Request) {
	var novaMusica model.Musica
	err := json.NewDecoder(r.Body).Decode(&novaMusica)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	novaMusica.ID = uuid.New().String()
	musicas = append(musicas, novaMusica)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(novaMusica)
} //criar

func GetMusica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, musica := range musicas {
		if musica.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(musica)
			return
		}
	}
	http.Error(w, "Musica nao encontrada", http.StatusNotFound)
}
