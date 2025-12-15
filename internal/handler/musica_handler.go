package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/estrelandoana/api-golang-treino/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GinListarMusicas(c *gin.Context) {
	musicas, err := service.ListarMusicas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Data: nil,
			Erro: "Erro na busca",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Data: musicas,
		Erro: nil,
	})
}

func GinCreateMusica(c *gin.Context) {
	var body entity.Musica
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Data: nil,
			Erro: "Dados inválidos",
		})
		return
	}
	if body.Titulo == "" || body.Artista == "" || body.Album == "" ||
		body.Ano <= 1000 || body.Ano > 2100 || body.Genero == "" ||
		body.Duracao <= 0 {
		c.JSON(http.StatusBadRequest, Response{
			Data: nil,
			Erro: "Campo obrigatório",
		})
		return
	}
	err = service.CreateMusica(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Data: nil,
			Erro: "Erro ao criar musica",
		})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Data: body,
		Erro: nil,
	})
}

func GinGetMusica(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Data: nil,
			Erro: "ID inválido",
		})
		return
	}
	musica, err := service.GetMusica(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Data: nil,
			Erro: "Música não encontrada",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Data: musica,
		Erro: nil,
	})
}

func GinUpdateMusica(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Data: nil,
			Erro: "ID inválido",
		})
		return
	}
	var body entity.Musica
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Data: nil,
			Erro: "Dados inválidos",
		})
		return
	}
	if body.Titulo == "" || body.Artista == "" || body.Album == "" ||
		body.Ano <= 1000 || body.Ano > 2100 || body.Genero == "" ||
		body.Duracao <= 0 {
		c.JSON(http.StatusBadRequest, Response{
			Data: nil,
			Erro: "Campo obrigatório",
		})
		return
	}
	musica, err := service.UpdateMusica(uint(idUint), body)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, Response{
				Data: nil,
				Erro: "Música não encontrada",
			})
		} else {
			c.JSON(http.StatusInternalServerError, Response{
				Data: nil,
				Erro: "Erro na atualização",
			})
		}
		return
	}
	c.JSON(http.StatusOK, Response{
		Data: musica,
		Erro: nil,
	})
}

func GinDeleteMusica(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Data: nil,
			Erro: "ID inválido",
		})
		return
	}
	err = service.DeleteMusica(uint(idUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, Response{
				Data: nil,
				Erro: "Música nao encontrada",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, Response{
				Data: nil,
				Erro: "Erro ao deletar musica",
			})
		}
		return
	}
	c.Status(http.StatusNoContent)
	return
}
