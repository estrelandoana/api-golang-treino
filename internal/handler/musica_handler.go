package handler

import (
	"encoding/json"
	"net/http"

	"github.com/estrelandoana/api-golang-treino/internal/db"
	entity "github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ListarMusicas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var musicas []entity.Musica
	result := db.DB.Find(&musicas)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Erro na busca de músicas"})
		return
	}
	json.NewEncoder(w).Encode(musicas)
} //listar

func CreateMusica(w http.ResponseWriter, r *http.Request) {
	var novaMusica entity.Musica
	err := json.NewDecoder(r.Body).Decode(&novaMusica)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Dados inválidos"})
		return
	}
	if novaMusica.Titulo == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Título obrigatório"})
		return
	}
	if novaMusica.Artista == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Artista obrigatório"})
		return
	}
	if novaMusica.Album == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Álbum obrigatório"})
		return
	}
	if novaMusica.Ano <= 1000 || novaMusica.Ano > 2100 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Ano inválido"})
		return
	}
	if novaMusica.Genero == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Gênero obrigatório"})
		return
	}
	if novaMusica.Duracao <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Duração inválida"})
		return
	}
	result := db.DB.Create(&novaMusica)
	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Erro ao criar música"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novaMusica)
} //criar

func GetMusica(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var musica entity.Musica
	result := db.DB.First(&musica, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"erro": "Música não encontrada"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"erro": "Erro na busca"})
		}
		return
	}
	json.NewEncoder(w).Encode(musica)
}

func UpdateMusica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var musica entity.Musica
	result := db.DB.First(&musica, id)
	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Música não encontrada"})
		return
	}
	var body entity.Musica
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Dados inválidos"})
		return
	}
	if body.Titulo == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Título obrigatório"})
		return
	}
	if body.Artista == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Artista obrigatório"})
		return
	}
	if body.Album == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Álbum obrigatório"})
		return
	}
	if body.Ano <= 1000 || body.Ano > 2100 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Ano inválido"})
		return
	}
	if body.Genero == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Gênero obrigatório"})
		return
	}
	if body.Duracao <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Duração inválida"})
		return
	}
	musica.Titulo = body.Titulo
	musica.Artista = body.Artista
	musica.Album = body.Album
	musica.Ano = body.Ano
	musica.Genero = body.Genero
	musica.Duracao = body.Duracao
	result = db.DB.Save(&musica)
	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Erro na atualização"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(musica)
}

/*************  ✨ Windsurf Command ⭐  *************/
// DeleteMusica - deleta uma música com base no seu ID
// Exemplo de resposta:
// {
// 	"erro": "Música não encontrada"
// }
/*******  cdb86c4c-0f60-4d84-9859-09b727f10446  *******/
func DeleteMusica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var musica entity.Musica
	result := db.DB.First(&musica, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"erro": "Música não encontrada"})
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"erro": "Erro na busca"})
			return
		}
	}
	result = db.DB.Delete(&musica)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Erro ao deletar"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
