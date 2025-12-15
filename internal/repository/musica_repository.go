package repository

import (
	"github.com/estrelandoana/api-golang-treino/internal/db"
	"github.com/estrelandoana/api-golang-treino/internal/entity"
	"gorm.io/gorm"
)

func ListarMusicas() ([]entity.Musica, error) {
	var musicas []entity.Musica
	err := db.DB.Find(&musicas).Error
	return musicas, err
}

func CreateMusica(musica entity.Musica) error {
	return db.DB.Create(&musica).Error
}

func GetMusica(id uint) (entity.Musica, error) {
	var musica entity.Musica
	err := db.DB.First(&musica, id).Error
	return musica, err
}

func UpdateMusica(musica entity.Musica) error {
	return db.DB.Save(&musica).Error
}

func DeleteMusica(id uint) error {
	result := db.DB.Delete(&entity.Musica{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
