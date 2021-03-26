package service

import (
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/repository"
)

type MangaMethods interface {
	AddManga(name string) (int, error)
	DeleteManga(id int) error
	GetManga(id int) (*models.Manga, error)
	GetMangas(page int) ([]*models.Manga, error)
}

type Service struct {
	MangaMethods
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		MangaMethods: NewMangaService(r),
	}
}
