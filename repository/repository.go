package repository

import (
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/jmoiron/sqlx"
)

type MangaRepository interface {
	GetManga(id int) (models.Manga, error)
	DeleteManga(id int) (bool, error)
	AddManga(manga models.Manga) (int, error)
	GetMangas() ([]models.Manga, error)
}

type Repository struct {
	MangaRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MangaRepository: NewMangaRepositoryPostgres(db),
	}
}
