package controller

import (
	"net/http"

	"github.com/estrelandoana/api-golang-treino/internal/dto"
	"github.com/estrelandoana/api-golang-treino/internal/service"
	"github.com/gin-gonic/gin"
)

// GinListarMusicas godoc
// @Summary      Lista todas as músicas
// @Description  Retorna todas as músicas cadastradas
// @Tags         musicas
// @Produce      json
// @Success      200  {object}  handler.Response
// @Failure      500  {object}  handler.Response
// @Router       /musicas [get]
func ListarMusica(c *gin.Context) {
	musicas, err := service.ListarMusicas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro na busca"})
		return
	}
	resp := make([]dto.MusicaResponse, 0, len(musicas))
	for _, m := range musicas {
		resp = append(resp, dto.MusicaResponse(m))
	}
	c.JSON(http.StatusOK, resp)
}
