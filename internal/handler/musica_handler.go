package handler

import (
	"encoding/json"
	"net/http"

	entity "github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var musicas = []entity.Musica{
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
	var novaMusica entity.Musica
	err := json.NewDecoder(r.Body).Decode(&novaMusica)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	if novaMusica.Titulo == "" {
		http.Error(w, "Título obrigatório", http.StatusBadRequest)
		return
	}
	if novaMusica.Artista == "" {
		http.Error(w, "Artista obrigatório", http.StatusBadRequest)
		return
	}
	if novaMusica.Album == "" {
		http.Error(w, "Álbum obrigatório", http.StatusBadRequest)
		return
	}
	if novaMusica.Ano <= 1000 || novaMusica.Ano > 2100 {
		http.Error(w, "Ano inválido", http.StatusBadRequest)
		return
	}
	if novaMusica.Genero == "" {
		http.Error(w, "Gênero obrigatório", http.StatusBadRequest)
		return
	}
	if novaMusica.Duracao <= 0 {
		http.Error(w, "Duração inválida", http.StatusBadRequest)
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

func UpdateMusica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var musicaAtualizada entity.Musica
	err := json.NewDecoder(r.Body).Decode(&musicaAtualizada)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	if musicaAtualizada.Titulo == "" {
		http.Error(w, "Título obrigatório", http.StatusBadRequest)
		return
	}
	if musicaAtualizada.Artista == "" {
		http.Error(w, "Artista obrigatório", http.StatusBadRequest)
		return
	}
	if musicaAtualizada.Album == "" {
		http.Error(w, "Álbum obrigatório", http.StatusBadRequest)
		return
	}
	if musicaAtualizada.Ano <= 1000 || musicaAtualizada.Ano > 2100 {
		http.Error(w, "Ano inválido", http.StatusBadRequest)
		return
	}
	if musicaAtualizada.Genero == "" {
		http.Error(w, "Gênero obrigatório", http.StatusBadRequest)
		return
	}
	if musicaAtualizada.Duracao <= 0 {
		http.Error(w, "Duração inválida", http.StatusBadRequest)
		return
	}
	for i, musica := range musicas {
		if musica.ID == id {
			musicaAtualizada.ID = id
			musicas[i] = musicaAtualizada
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(musicaAtualizada)
			return
		}
	}
	http.Error(w, "Música não encontrada", http.StatusNotFound)
}

func DeleteMusica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, musica := range musicas {
		if musica.ID == id {
			musicas = append(musicas[:i], musicas[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Música não encontrada", http.StatusNotFound)
}
