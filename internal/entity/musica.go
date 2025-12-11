package entity

type Musica struct {
	ID      uint   `gorm:"primaryKey"`
	Titulo  string `json:"titulo"`
	Artista string `json:"artista"`
	Album   string `json:"album"`
	Ano     int    `json:"ano"`
	Genero  string `json:"genero"`
	Duracao int    `json:"duracao"`
}
