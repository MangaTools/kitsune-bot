package repository

import (
	"errors"
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/sirupsen/logrus"
)

type ChapterRepositoryPostgres struct {
	db DbWorker
}

func (c ChapterRepositoryPostgres) GetChapter(chapterId int) (*models.Chapter, error) {
	chapter := new(models.Chapter)
	query := fmt.Sprintf("SELECT id, manga_id, number, pages FROM %s WHERE id = $1", chapterTable)
	err := c.db.QueryRow(query, chapterId).Scan(&chapter.Id, &chapter.MangaId, &chapter.Number, &chapter.Pages)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("Не удалось получить главу")
	}
	return chapter, nil
}

func (c ChapterRepositoryPostgres) CreateChapter(mangaId int, chapter float32, pages int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (manga_id, number, pages, status) VALUES($1, $2, $3, $4) RETURNING id", chapterTable)
	err := c.db.QueryRow(query, mangaId, chapter, pages, models.InWorkChapter).Scan(&id)
	if err != nil {
		logrus.Error(err)
		return -1, errors.New("Не удалось создать главу.")
	}
	return id, nil
}

func (c ChapterRepositoryPostgres) DeleteChapter(chapterId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", chapterTable)
	_, err := c.db.Exec(query, chapterId)
	if err != nil {
		logrus.Error(err)
		return errors.New("Ошибка при удалении главы.")
	}
	return nil
}

func (c ChapterRepositoryPostgres) HasChapter(chapterId int) bool {
	var hasChapter bool
	query := fmt.Sprintf("SELECT exists(SELECT * FROM %s WHERE id=$1)", chapterTable)
	err := c.db.QueryRow(query, chapterId).Scan(&hasChapter)
	if err != nil {
		logrus.Error(err)
		return false
	}
	return hasChapter
}

func NewChapterRepositoryPostgres(db DbWorker) *ChapterRepositoryPostgres {
	return &ChapterRepositoryPostgres{db: db}
}
