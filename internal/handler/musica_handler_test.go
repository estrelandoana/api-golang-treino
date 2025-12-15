package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/estrelandoana/api-golang-treino/internal/db"
	entity "github.com/estrelandoana/api-golang-treino/internal/entity"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/musicas", GinListarMusicas)
		api.POST("/musicas", GinCreateMusica)
		api.GET("/musicas/:id", GinGetMusica)
		api.PUT("/musicas/:id", GinUpdateMusica)
		api.DELETE("/musicas/:id", GinDeleteMusica)
	}
	return r
}
func TestMain(m *testing.M) {
	db.ConectorDB()
	db.DB.AutoMigrate(&entity.Musica{})
	code := m.Run()
	os.Exit(code)
}

func TestListarMusicas(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db.DB.Exec("DELETE FROM musicas")
	db.DB.Create(&entity.Musica{
		Titulo:  "Nome da musica",
		Artista: "Artista da musica",
		Album:   "Album da musica",
		Ano:     2022,
		Genero:  "Genero da musica",
		Duracao: 30,
	})
	router := setupRouter()
	req, err := http.NewRequest("GET", "/api/v1/musicas", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expectativa: %v, obtido: %v", http.StatusOK, w.Code)
	}
	if w.Body.Len() == 0 {
		t.Errorf("Expectativa: corpo com música, recebido: vazio")
	}
}

func TestCreateMusica(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db.DB.Exec("DELETE FROM musicas")
	router := setupRouter()
	body := `{"titulo":"TituloTeste","artista":"ArtistaTeste","album":"AlbumTeste","ano":2017,"genero":"GeneroTeste","duracao":8}`
	req, err := http.NewRequest("POST", "/api/v1/musicas", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expectativa: %v ou %v, obtida: %v", http.StatusOK, http.StatusCreated, w.Code)
	}
	if w.Body.Len() == 0 {
		t.Errorf("Expectativa: corpo com música, recebido: vazio")
	}
}

func TestGetMusica(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db.DB.Exec("DELETE FROM musicas")
	musica := entity.Musica{
		Titulo:  "Nome da musica",
		Artista: "Artista da musica",
		Album:   "Album da musica",
		Ano:     2022,
		Genero:  "Genero da musica",
		Duracao: 30,
	}
	db.DB.Create(&musica)
	router := setupRouter()
	req, err := http.NewRequest("GET", "/api/v1/musicas/"+fmt.Sprint(musica.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expectativa: status 200, obtido: %v", w.Code)
	}
}

func TestUpdateMusica(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db.DB.Exec("DELETE FROM musicas")
	musica := entity.Musica{
		Titulo:  "Nome da musica",
		Artista: "Artista da musica",
		Album:   "Album da musica",
		Ano:     2022,
		Genero:  "Genero da musica",
		Duracao: 30,
	}
	db.DB.Create(&musica)
	body := `{"titulo":"TituloTeste","artista":"ArtistaTeste","album":"AlbumTeste","ano":2017,"genero":"GeneroTeste","duracao":8}`
	router := setupRouter()
	req, err := http.NewRequest("PUT", "/api/v1/musicas/"+fmt.Sprint(musica.ID), strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expectativa: status 200, obtido: %v", w.Code)
	}
}

func TestDeleteMusica(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db.DB.Exec("DELETE FROM musicas")
	musica := entity.Musica{
		Titulo:  "Nome da musica",
		Artista: "Artista da musica",
		Album:   "Album da musica",
		Ano:     2022,
		Genero:  "Genero da musica",
		Duracao: 30,
	}
	db.DB.Create(&musica)
	router := setupRouter()
	req, err := http.NewRequest("DELETE", "/api/v1/musicas/"+fmt.Sprint(musica.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("Expectativa: status 204 No Content, obtido: %v", w.Code)
	}
}
