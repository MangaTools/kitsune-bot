package repository

import (
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/jmoiron/sqlx"
)

type MangaRepository interface {
	CreateManga(name string, status models.MangaStatus) (int, error)
	DeleteManga(id int) error
	HasManga(id int) bool
	GetManga(id int) (*models.Manga, error)
	GetMangas(max int, page int) ([]*models.Manga, error)
}

type Repository struct {
	MangaRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MangaRepository: NewMangaRepositoryPostgres(db),
	}
}
