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

type ChapterRepository interface {
	CreateChapter(mangaId int, chapter float32, pages int) (int, error)
	DeleteChapter(chapterId int) error
	HasChapter(chapterId int) bool
	GetChapter(chapterId int) (*models.Chapter, error)
}

type UserRepository interface {
	CreateUser(userId string, username string) (*models.User, error)
	GetUser(userId string) (*models.User, error)
	HasUser(userId string) bool
	GetTopUsers(characteristic string) ([]*models.User, error)
	AddToField(userid string, characteristic models.UserCharacteristic, value int) (int, error)
}

type WorkRepository interface {
	GetWork(workId int) (*models.Owner, error)
	GetWorksByWorkType(chapterId int, workType models.WorkType) ([]*models.Owner, error)
	CreateWork(userId string, chapterId int, pageStart int, pageEnd int, workType models.WorkType) (int, error)
	SetWorkStatus(workId int, status models.OwnerPageStatus) error
	DeleteWork(workId int) error
	HasWork(workId int) bool
	MergeWorks([][]*models.Owner) error
}

type Repository struct {
	MangaRepository
	ChapterRepository
	UserRepository
	WorkRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MangaRepository:   NewMangaRepositoryPostgres(db),
		ChapterRepository: NewChapterRepositoryPostgres(db),
		UserRepository:    NewUserRepositoryPostgres(db),
		WorkRepository:    NewWorkRepositoryPostgres(db),
	}
}
