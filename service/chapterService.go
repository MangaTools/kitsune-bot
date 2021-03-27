package service

import (
	"errors"
	"github.com/ShaDream/kitsune-bot/repository"
)

type ChapterService struct {
	chapterRepo repository.ChapterRepository
	mangaRepo   repository.MangaRepository
}

func NewChapterService(chapterRepo repository.ChapterRepository, mangaRepo repository.MangaRepository) *ChapterService {
	return &ChapterService{chapterRepo: chapterRepo, mangaRepo: mangaRepo}
}

func (c ChapterService) AddChapter(mangaId int, chapter float32, pages int) (int, error) {
	if c.mangaRepo.HasManga(mangaId) {
		return -1, errors.New("Манги с таким ID не существует.")
	}
	chapterId, err := c.chapterRepo.CreateChapter(mangaId, chapter, pages)
	if err != nil {
		return -1, err
	}
	return chapterId, nil
}

func (c ChapterService) DeleteChapter(chapterId int) error {
	if !c.chapterRepo.HasChapter(chapterId) {
		return errors.New("Главы с таким ID не существует.")
	}
	err := c.chapterRepo.DeleteChapter(chapterId)
	if err != nil {
		return err
	}

	return nil
}
