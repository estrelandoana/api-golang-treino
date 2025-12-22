package controller

import (
	"net/http"
	"strconv"

	"github.com/estrelandoana/api-golang-treino/internal/dto"
	"github.com/estrelandoana/api-golang-treino/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateMusica(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	var req dto.MusicaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}
	m, err := service.GetMusica(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"erro": "Música não encontrada"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro na busca"})
		}
		return
	}
	m.Titulo = req.Titulo
	m.Artista = req.Artista
	m.Album = req.Album
	m.Ano = req.Ano
	m.Genero = req.Genero
	m.Duracao = req.Duracao

	m, err = service.UpdateMusica(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro na atualização"})
		return
	}
	resp := dto.MusicaResponse(m)
	c.JSON(http.StatusOK, resp)
}
