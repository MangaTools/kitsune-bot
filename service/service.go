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

type UserMethods interface {
	GetUser(userId string, username string) (*models.User, error)
	GetTopUser(characteristic models.UserCharacteristic) ([]*models.User, error)
	TryCreateUser(userId string, username string) error
}

type WorkMethods interface {
	CheckWork(workId int) error
	BookWork(userId string, chapterId int, workType models.WorkType, startPage int, endPage int) (int, error)
	RemoveBookedWork(workId int) error
	DoneWork(workId int) error
}

type AccessMethods interface {
	SetRoleAccess(roleId string, access models.RoleAccess) error
	RemoveRoleAccess(roleId string) error
	HasRoleAccess(roleId string) bool
	GetAllRoleAccesses() ([]*models.Role, error)
}

type Service struct {
	MangaMethods
	ChapterMethods
	UserMethods
	WorkMethods
	AccessMethods
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		MangaMethods:   NewMangaService(*r),
		ChapterMethods: NewChapterService(*r),
		UserMethods:    NewUserService(*r),
		WorkMethods:    NewWorkService(*r),
		AccessMethods:  NewAccessService(*r),
	}
}
