package model

type Musica struct {
	ID      string `json:"id"`
	Titulo  string `json:"titulo"`
	Artista string `json:"artista"`
	Album   string `json:"album"`
	Ano     int    `json:"ano"`
	Genero  string `json:"genero"`
	Duracao int    `json:"duracao"`
}
