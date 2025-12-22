package controller

import (
	"net/http"
	"strconv"

	"github.com/estrelandoana/api-golang-treino/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GinDeleteMusica godoc
// @Summary      Remove uma música
// @Description  Deleta uma música pelo ID
// @Tags         musicas
// @Produce      json
// @Param        id   path  int  true  "ID da música"
// @Success      204  {object}  nil
// @Failure      400  {object}  handler.Response
// @Failure      404  {object}  handler.Response
// @Failure      500  {object}  handler.Response
// @Router       /musicas/{id} [delete]
func DeleteMusica(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	err = service.DeleteMusica(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"erro": "Música não encontrada"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao deletar música"})
		}
		return
	}
	c.Status(http.StatusNoContent)
}
