package service

import (
	"errors"
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/repository"
	"github.com/sirupsen/logrus"
)

type MangaService struct {
	repo repository.Repository
}

func (m MangaService) AddManga(name string) (int, error) {
	id, err := m.repo.MangaRepository.CreateManga(name, models.TranslatingManga)
	if err != nil {
		err = errors.New(fmt.Sprintf("Что-то пошло не так при добавлении манги с именем %s, ошибка %s", name, err))
		logrus.Error(err)
		return -1, err
	}
	return id, nil
}

func (m MangaService) DeleteManga(id int) error {
	if m.HasManga(id) {
		return m.repo.MangaRepository.DeleteManga(id)
	}
	return errors.New("Манги с таким ChapterId не существует.")
}

func (m MangaService) HasManga(id int) bool {
	return m.repo.MangaRepository.HasManga(id)
}

func (m MangaService) GetManga(id int) (*models.Manga, error) {
	if m.HasManga(id) {
		return m.repo.MangaRepository.GetManga(id)
	}
	return nil, errors.New("Манги с таким ChapterId не существует.")
}

func (m MangaService) GetMangas(page int) ([]*models.Manga, error) {
	return m.repo.MangaRepository.GetMangas(10, page)
}

func NewMangaService(repo repository.Repository) *MangaService {
	return &MangaService{repo: repo}
}
