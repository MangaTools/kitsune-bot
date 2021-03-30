package service

import (
	"errors"
	"github.com/ShaDream/kitsune-bot/repository"
)

type ChapterService struct {
	repo repository.Repository
}

func NewChapterService(chapterRepo repository.Repository) *ChapterService {
	return &ChapterService{chapterRepo}
}

func (c ChapterService) AddChapter(mangaId int, chapter float32, pages int) (int, error) {
	if !c.repo.MangaRepository.HasManga(mangaId) {
		return -1, errors.New("Манги с таким ID не существует.")
	}
	chapterId, err := c.repo.ChapterRepository.CreateChapter(mangaId, chapter, pages)
	if err != nil {
		return -1, err
	}
	return chapterId, nil
}

func (c ChapterService) DeleteChapter(chapterId int) error {
	if !c.repo.ChapterRepository.HasChapter(chapterId) {
		return errors.New("Главы с таким ID не существует.")
	}
	err := c.repo.ChapterRepository.DeleteChapter(chapterId)
	if err != nil {
		return err
	}

	return nil
}
