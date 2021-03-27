package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ChapterRepositoryPostgres struct {
	db *sqlx.DB
}

func (c ChapterRepositoryPostgres) CreateChapter(mangaId int, chapter float32, pages int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (manga_id, number, pages) VALUES($1, $2, $3) RETURNING id", chapterTable)
	err := c.db.QueryRow(query, mangaId, chapter, pages).Scan(&id)
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

func NewChapterRepositoryPostgres(db *sqlx.DB) *ChapterRepositoryPostgres {
	return &ChapterRepositoryPostgres{db: db}
}
