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
	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	db.ConectorDB()
	db.DB.AutoMigrate(&entity.Musica{})
	code := m.Run()
	os.Exit(code)
}

func TestListarMusicas(t *testing.T) {
	db.DB.Exec("DELETE FROM musicas")
	db.DB.Create(&entity.Musica{
		Titulo:  "Nome da musica",
		Artista: "Artista da musica",
		Album:   "Album da musica",
		Ano:     2022,
		Genero:  "Genero da musica",
		Duracao: 30,
	})
	req, err := http.NewRequest("GET", "/musicas", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListarMusicas)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expectativa: %v, obtido: %v", http.StatusOK, status)
	}
	if rr.Body.Len() == 0 {
		t.Errorf("Expectativa: corpo com música, recebido: vazio")
	}
}

func TestCreateMusica(t *testing.T) {
	db.DB.Exec("DELETE FROM musicas")
	body := `{"titulo":"TituloTeste","artista":"ArtistaTeste","album":"AlbumTeste","ano":2017,"genero":"GeneroTeste","duracao":8}`
	req, err := http.NewRequest("POST", "/musicas", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMusica)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK && status != http.StatusCreated {
		t.Errorf("Expectativa: %v ou %v, obtida: %v", http.StatusOK, http.StatusCreated, status)
	}
	if rr.Body.Len() == 0 {
		t.Errorf("Expectativa: corpo com música, recebido: vazio")
	}
}

func TestGetMusica(t *testing.T) {
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
	req, err := http.NewRequest("GET", "/musicas/"+fmt.Sprint(musica.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/musicas/{id}", GinGetMusica).Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expectativa: status 200, obtido: %v", status)
	}
}

func TestUpdateMusica(t *testing.T) {
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
	req, err := http.NewRequest("PUT", "/musicas/"+fmt.Sprint(musica.ID), strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/musicas/{id}", UpdateMusica).Methods("PUT")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expectativa: status 200, obtido: %v", status)
	}
}

func TestDeleteMusica(t *testing.T) {
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
	req, err := http.NewRequest("DELETE", "/musicas/"+fmt.Sprint(musica.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/musicas/{id}", DeleteMusica).Methods("DELETE")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Expectativa: status 204 No Content, obtido: %v", status)
	}
}
