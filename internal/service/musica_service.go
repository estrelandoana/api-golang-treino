package service

import (
	"github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/estrelandoana/api-golang-treino/internal/repository"
)

func CreateMusica(m entity.Musica) (entity.Musica, error) {
	return repository.CreateMusica(m)
}

func ListarMusicas() ([]entity.Musica, error) {
	return repository.ListarMusicas()
}

func GetMusica(id uint) (entity.Musica, error) {
	return repository.GetMusica(id)
}

func UpdateMusica(m entity.Musica) (entity.Musica, error) {
	return repository.UpdateMusica(m)
}

func DeleteMusica(id uint) error {
	return repository.DeleteMusica(id)
}
