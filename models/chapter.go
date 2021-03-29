package models

import (
	"fmt"
)

type ChapterStatus int

const (
	DoneChapter ChapterStatus = iota
	InWorkChapter
)

var chapterStatusesToString = map[ChapterStatus]string{
	DoneChapter:   "Готова",
	InWorkChapter: "В процессе",
}

func IsValidChapterStatus(id ChapterStatus) bool {
	switch id {
	case DoneChapter, InWorkChapter:
		return true
	}
	return false
}

func GetChapterStatusString(id ChapterStatus) string {
	if val, ok := chapterStatusesToString[id]; ok {
		return fmt.Sprintf("%d - %s", int(id), val)
	}
	return "Такого статуса нет"
}

func GetAllChapterStatusesString() string {
	result := ""
	for id, val := range chapterStatusesToString {
		result += fmt.Sprintf("%d - %s\n", id, val)
	}
	return result
}

type Chapter struct {
	Id         int           `json:"id" db:"id"`
	MangaId    int           `json:"manga_id" db:"manga_id"`
	Number     float32       `json:"number" db:"number"`
	Pages      int           `json:"pages" db:"pages"`
	CleanOwner []Owner       `json:"clean_owner" db:"-"`
	TyperOwner []Owner       `json:"typer_owner" db:"-"`
	Translator []Owner       `json:"translator" db:"-"`
	Editor     []Owner       `json:"editor" db:"-"`
	Status     ChapterStatus `json:"status" db:"status"`
}
