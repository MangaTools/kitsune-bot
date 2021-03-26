package repository

import (
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/jmoiron/sqlx"
)

type MangaRepositoryPostgres struct {
	db *sqlx.DB
}

func (m MangaRepositoryPostgres) GetManga(id int) (models.Manga, error) {
	panic("implement me")
}

func (m MangaRepositoryPostgres) DeleteManga(id int) (bool, error) {
	panic("implement me")
}

func (m MangaRepositoryPostgres) AddManga(manga models.Manga) (int, error) {
	panic("implement me")
}

func (m MangaRepositoryPostgres) GetMangas() ([]models.Manga, error) {
	panic("implement me")
}

func NewMangaRepositoryPostgres(db *sqlx.DB) *MangaRepositoryPostgres {
	return &MangaRepositoryPostgres{db: db}
}
