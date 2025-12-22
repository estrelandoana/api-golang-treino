package controller

import (
	"net/http"

	"github.com/estrelandoana/api-golang-treino/internal/dto"
	"github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/estrelandoana/api-golang-treino/internal/service"
	"github.com/gin-gonic/gin"
)

// GinCreateMusica godoc
// @Summary      Cria nova música
// @Description  Criação de uma nova música
// @Tags         musicas
// @Accept       json
// @Produce      json
// @Param        musica  body  entity.Musica  true  "Música para criar"
// @Success      201  {object}  handler.Response
// @Failure      400  {object}  handler.Response
// @Failure      500  {object}  handler.Response
// @Router       /musicas [post]
func CreateMusica(c *gin.Context) {
	var req dto.MusicaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}
	musica := entity.Musica{
		Titulo:  req.Titulo,
		Artista: req.Artista,
		Album:   req.Album,
		Ano:     req.Ano,
		Genero:  req.Genero,
		Duracao: req.Duracao,
	}
	mCriado, err := service.CreateMusica(musica)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao criar musica"})
		return
	}
	resp := dto.MusicaResponse{
		ID:      mCriado.ID,
		Titulo:  mCriado.Titulo,
		Artista: mCriado.Artista,
		Album:   mCriado.Album,
		Ano:     mCriado.Ano,
		Genero:  mCriado.Genero,
		Duracao: mCriado.Duracao,
	}
	c.JSON(http.StatusCreated, resp)
}
