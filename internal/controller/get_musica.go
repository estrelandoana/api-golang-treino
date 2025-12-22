package controller

import (
	"net/http"
	"strconv"

	"github.com/estrelandoana/api-golang-treino/internal/dto"
	"github.com/estrelandoana/api-golang-treino/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GinGetMusica godoc
// @Summary      Busca uma música por ID
// @Description  Retorna uma música específica pelo ID
// @Tags         musicas
// @Produce      json
// @Param        id   path      int  true  "ID da música"
// @Success      200  {object}  handler.Response
// @Failure      400  {object}  handler.Response
// @Failure      404  {object}  handler.Response
// @Router       /musicas/{id} [get]
func GetMusica(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
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
	resp := dto.MusicaResponse(m)
	c.JSON(http.StatusOK, resp)
}
