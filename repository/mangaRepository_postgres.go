package repository

import (
	"errors"
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MangaRepositoryPostgres struct {
	db *sqlx.DB
}

func (m MangaRepositoryPostgres) CreateManga(name string, status models.MangaStatus) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, status) VALUES($1, $2) RETURNING id", mangaTable)
	err := m.db.QueryRow(query, name, status).Scan(&id)
	if err != nil {
		logrus.Error(err)
		return -1, errors.New("Не удалось создать мангу.")
	}
	return id, nil
}

func (m MangaRepositoryPostgres) DeleteManga(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", mangaTable)
	_, err := m.db.Exec(query, id)
	if err != nil {
		logrus.Error(err)
		return errors.New("Ошибка при удалении манги.")
	}
	return nil
}

func (m MangaRepositoryPostgres) HasManga(id int) bool {
	var hasManga bool
	query := fmt.Sprintf("SELECT exists(SELECT * FROM %s WHERE id=$1)", mangaTable)
	err := m.db.QueryRow(query, id).Scan(&hasManga)
	if err != nil {
		logrus.Error(err)
		return false
	}
	return hasManga
}

func (m MangaRepositoryPostgres) GetManga(id int) (*models.Manga, error) {
	manga := new(models.Manga)
	query := fmt.Sprintf("SELECT id, name, status FROM %s WHERE id=$1", mangaTable)
	err := m.db.QueryRow(query, id).Scan(&manga.Id, &manga.Name, &manga.Status)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("Произошла ошибка при получении манги.")
	}
	return manga, nil
}

func (m MangaRepositoryPostgres) GetMangas(max int, page int) ([]*models.Manga, error) {
	var mangas = make([]*models.Manga, 0)
	query := fmt.Sprintf("SELECT id, name, status FROM %s LIMIT $1 OFFSET $2", mangaTable)
	rows, err := m.db.Query(query, max, (page-1)*max)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("Произошла ошибка при получении манги.")
	}
	for rows.Next() {
		manga := new(models.Manga)
		err := rows.Scan(&manga.Id, &manga.Name, &manga.Status)
		if err != nil {
			logrus.Error(err)
			return nil, errors.New("Произошла ошибка при получении манги.")
		}
		mangas = append(mangas, manga)
	}
	return mangas, nil
}

func NewMangaRepositoryPostgres(db *sqlx.DB) *MangaRepositoryPostgres {
	return &MangaRepositoryPostgres{db: db}
}
