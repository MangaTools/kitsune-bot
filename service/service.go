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

type ChapterMethods interface {
	AddChapter(mangaId int, chapter float32, pages int) (int, error)
	DeleteChapter(chapterId int) error
}

type Service struct {
	MangaMethods
	ChapterMethods
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		MangaMethods:   NewMangaService(r.MangaRepository),
		ChapterMethods: NewChapterService(r.ChapterRepository, r.MangaRepository),
	}
}
