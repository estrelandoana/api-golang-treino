package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/estrelandoana/api-golang-treino/internal/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ListarMusicas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	musicas, err := service.ListarMusicas()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "Erro na busca",
		})
		return
	}
	json.NewEncoder(w).Encode(Response{
		Data: musicas,
		Erro: nil,
	})
}

func CreateMusica(w http.ResponseWriter, r *http.Request) {
	var novaMusica entity.Musica
	err := json.NewDecoder(r.Body).Decode(&novaMusica)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "Dados inválidos",
		})
		return
	}
	if novaMusica.Titulo == "" || novaMusica.Artista == "" || novaMusica.Album == "" ||
		novaMusica.Ano <= 1000 || novaMusica.Ano > 2100 || novaMusica.Genero == "" ||
		novaMusica.Duracao <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "Campo obrigatório",
		})
		return
	}
	err = service.CreateMusica(novaMusica)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "Erro ao criar musica",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Data: novaMusica,
		Erro: nil,
	})
}

func GetMusica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "ID inválido",
		})
		return
	}
	musica, err := service.GetMusica(uint(idUint))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "Música não encontrada",
		})
		return
	}
	json.NewEncoder(w).Encode(Response{
		Data: musica,
		Erro: nil,
	})
}

func UpdateMusica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "ID inválido",
		})
		return
	}
	var body entity.Musica
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "Dados inválidos",
		})
		return
	}
	if body.Titulo == "" || body.Artista == "" || body.Album == "" ||
		body.Ano <= 1000 || body.Ano > 2100 || body.Genero == "" ||
		body.Duracao <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "Campo obrigatório",
		})
		return
	}
	musica, err := service.UpdateMusica(uint(idUint), body)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Data: nil,
				Erro: "Música não encontrada",
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Data: nil,
				Erro: "Erro na atualização",
			})
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Data: musica,
		Erro: nil,
	})
}

func DeleteMusica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Data: nil,
			Erro: "ID inválido",
		})
		return
	}
	err = service.DeleteMusica(uint(idUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Data: nil,
				Erro: "Música nao encontrada",
			})
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Data: nil,
				Erro: "Erro ao deletar musica",
			})
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}
