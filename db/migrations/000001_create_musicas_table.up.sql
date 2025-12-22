CREATE TABLE musicas (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    artista VARCHAR(255) NOT NULL,
    album VARCHAR(255) NOT NULL,
    ano INT NOT NULL,
    genero VARCHAR(255) NOT NULL,
    duracao INT NOT NULL
)