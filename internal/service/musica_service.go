package service

import (
	"github.com/estrelandoana/api-golang-treino/internal/db"
	"github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/estrelandoana/api-golang-treino/internal/repository"
)

func ListarMusicas() ([]entity.Musica, error) {
	return repository.ListarMusicas()
}

func CreateMusica(musica entity.Musica) error {
	return repository.CreateMusica(musica)
}

func GetMusica(id uint) (entity.Musica, error) {
	return repository.GetMusica(id)
}

func UpdateMusica(id uint, body entity.Musica) (entity.Musica, error) {
	var musica entity.Musica
	err := db.DB.First(&musica, id).Error
	if err != nil {
		return musica, err
	}
	musica.Titulo = body.Titulo
	musica.Artista = body.Artista
	musica.Album = body.Album
	musica.Ano = body.Ano
	musica.Genero = body.Genero
	musica.Duracao = body.Duracao
	err = repository.UpdateMusica(musica)
	return musica, err
}

func DeleteMusica(id uint) error {
	return repository.DeleteMusica(id)
}
