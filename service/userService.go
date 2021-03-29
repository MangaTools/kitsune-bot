package service

import (
	"errors"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (u UserService) TryCreateUser(userId string, username string) error {
	if u.repo.HasUser(userId) {
		return nil
	}
	_, err := u.repo.CreateUser(userId, username)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) GetUser(userId string, username string) (*models.User, error) {
	var (
		user *models.User
		err  error
	)

	if !u.repo.HasUser(userId) {
		user, err = u.repo.CreateUser(userId, username)
	} else {
		user, err = u.repo.GetUser(userId)
	}

	if err != nil {
		return nil, errors.New("Не удалось показать юзера.")
	}
	return user, nil
}

var userCharacteristicToDBField = map[models.UserCharacteristic]string{
	models.UserCharacteristicScore:           "score",
	models.UserCharacteristicTranslatedPages: "translated_pages",
	models.UserCharacteristicEditedPages:     "edited_pages",
	models.UserCharacteristicCleanedPages:    "cleaned_pages",
	models.UserCharacteristicTypedPages:      "typed_chapters",
}

func (u UserService) GetTopUser(characteristic models.UserCharacteristic) ([]*models.User, error) {
	orderField := userCharacteristicToDBField[characteristic]
	users, err := u.repo.GetTopUsers(orderField)
	if err != nil {
		return nil, errors.New("Не удалось построить топ.")
	}
	return users, nil
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
