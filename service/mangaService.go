package service

import (
	"errors"
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/repository"
	"github.com/sirupsen/logrus"
)

type MangaService struct {
	repo repository.MangaRepository
}

func (m MangaService) AddManga(name string) (int, error) {
	id, err := m.repo.CreateManga(name, models.TranslatingManga)
	if err != nil {
		err = errors.New(fmt.Sprintf("Что-то пошло не так при добавлении манги с именем %s, ошибка %s", name, err))
		logrus.Error(err)
		return -1, err
	}
	return id, nil
}

func (m MangaService) DeleteManga(id int) error {
	if m.HasManga(id) {
		return m.repo.DeleteManga(id)
	}
	return errors.New("Манги с таким Id не существует.")
}

func (m MangaService) HasManga(id int) bool {
	return m.repo.HasManga(id)
}

func (m MangaService) GetManga(id int) (*models.Manga, error) {
	if m.HasManga(id) {
		return m.repo.GetManga(id)
	}
	return nil, errors.New("Манги с таким Id не существует.")
}

func (m MangaService) GetMangas(page int) ([]*models.Manga, error) {
	return m.repo.GetMangas(10, page)
}

func NewMangaService(repo repository.MangaRepository) *MangaService {
	return &MangaService{repo: repo}
}
