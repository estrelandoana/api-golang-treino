package repository

import (
	"github.com/estrelandoana/api-golang-treino/internal/database"
	"github.com/estrelandoana/api-golang-treino/internal/entity"
)

func CreateMusica(m entity.Musica) (entity.Musica, error) {
	err := database.DB.Create(&m).Error
	return m, err
}

func ListarMusicas() ([]entity.Musica, error) {
	var musicas []entity.Musica
	err := database.DB.Find(&musicas).Error
	return musicas, err
}

func GetMusica(id uint) (entity.Musica, error) {
	var m entity.Musica
	err := database.DB.First(&m, id).Error
	return m, err
}

func UpdateMusica(m entity.Musica) (entity.Musica, error) {
	err := database.DB.Save(&m).Error
	return m, err
}

func DeleteMusica(id uint) error {
	err := database.DB.Delete(&entity.Musica{}, id).Error
	return err
}
