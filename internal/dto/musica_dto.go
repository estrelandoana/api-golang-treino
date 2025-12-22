package dto

type MusicaRequest struct {
	Titulo  string `json:"titulo"`
	Artista string `json:"artista"`
	Album   string `json:"album"`
	Ano     int    `json:"ano"`
	Genero  string `json:"genero"`
	Duracao int    `json:"duracao"`
}

type MusicaResponse struct {
	ID      uint   `json:"id"`
	Titulo  string `json:"titulo"`
	Artista string `json:"artista"`
	Album   string `json:"album"`
	Ano     int    `json:"ano"`
	Genero  string `json:"genero"`
	Duracao int    `json:"duracao"`
}
