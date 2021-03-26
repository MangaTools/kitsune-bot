package models

import "fmt"

type MangaStatus int

const (
	DoneManga MangaStatus = iota
	TranslatingManga
)

var mangaStatusToString = map[MangaStatus]string{
	DoneManga:        "Готова",
	TranslatingManga: "Переводим",
}

func IsValidMangaStatus(id MangaStatus) bool {
	switch id {
	case DoneManga, TranslatingManga:
		return true
	}
	return false
}

func GetMangaStatusString(id MangaStatus) string {
	if val, ok := mangaStatusToString[id]; ok {
		return fmt.Sprintf("%d - %s", int(id), val)
	}
	return "Такого статуса нет"
}

func GetAllMangaStatusesString() string {
	result := ""
	for id, val := range mangaStatusToString {
		result += fmt.Sprintf("%d - %s\n", id, val)
	}
	return result
}

type Manga struct {
	Id       int         `json:"id" db:"id"`
	Name     string      `json:"name" db:"name"`
	Chapters []*Chapter  `json:"chapters" db:"-"`
	Status   MangaStatus `json:"status" db:"status"`
}

func NewManga(id int, name string) *Manga {
	return &Manga{Name: name, Id: id, Status: TranslatingManga, Chapters: make([]*Chapter, 0)}
}
