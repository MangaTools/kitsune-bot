package models

import (
	"fmt"
	"strings"
)

type UserCharacteristic int

const (
	UserCharacteristicScore UserCharacteristic = iota
	UserCharacteristicTranslatedPages
	UserCharacteristicEditedPages
	UserCharacteristicCheckedPages
	UserCharacteristicCleanedPages
	UserCharacteristicTypedPages
)

var userCharacteristicToString = map[UserCharacteristic]string{
	UserCharacteristicScore:           "Баллы",
	UserCharacteristicTranslatedPages: "Переведено",
	UserCharacteristicEditedPages:     "Отредактировано",
	UserCharacteristicCheckedPages:    "Проверено",
	UserCharacteristicCleanedPages:    "Отклинино",
	UserCharacteristicTypedPages:      "Затайплено",
}

func IsValidUserCharacteristic(id UserCharacteristic) bool {
	switch id {
	case UserCharacteristicScore,
		UserCharacteristicTranslatedPages,
		UserCharacteristicEditedPages,
		UserCharacteristicCheckedPages,
		UserCharacteristicCleanedPages,
		UserCharacteristicTypedPages:
		return true
	}
	return false
}

func GetUserCharacteristicString(id UserCharacteristic) string {
	if val, ok := userCharacteristicToString[id]; ok {
		return fmt.Sprintf("%s", val)
	}
	return "Такого статуса нет"
}

func GetAllUserCharacteristicsString() string {
	result := ""
	for id, val := range userCharacteristicToString {
		result += fmt.Sprintf("%d - %s\n", id, val)
	}
	return result
}

type User struct {
	Id              string `json:"id" db:"id"`
	Username        string
	Score           int `json:"score" db:"score"`
	TranslatedPages int `json:"translated_pages" db:"translated_pages"`
	EditedPages     int `json:"edited_pages" db:"edited_pages"`
	CheckedPages    int `json:"checked_pages" db:"checked_pages"`
	CleanedPages    int `json:"cleaned_pages" db:"cleaned_pages"`
	TypedPages      int `json:"typed_pages" db:"typed_pages"`
}

func (u User) GetInfo() string {
	characteristics := []string{
		fmt.Sprintf("%s - %d", userCharacteristicToString[UserCharacteristicScore], u.Score),
		fmt.Sprintf("%s - %d", userCharacteristicToString[UserCharacteristicTranslatedPages], u.TranslatedPages),
		fmt.Sprintf("%s - %d", userCharacteristicToString[UserCharacteristicEditedPages], u.EditedPages),
		fmt.Sprintf("%s - %d", userCharacteristicToString[UserCharacteristicCheckedPages], u.CheckedPages),
		fmt.Sprintf("%s - %d", userCharacteristicToString[UserCharacteristicCleanedPages], u.CleanedPages),
		fmt.Sprintf("%s - %d", userCharacteristicToString[UserCharacteristicTypedPages], u.TypedPages),
	}
	return strings.Join(characteristics, "\n")
}
