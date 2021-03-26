package models

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type ChapterStatus int

const (
	DoneChapter ChapterStatus = iota
	FindingPersonsChapter
	InWorkChapter
)

var chapterStatusesToString = map[ChapterStatus]string{
	DoneChapter:           "Готова",
	FindingPersonsChapter: "Ищем людей",
	InWorkChapter:         "В процессе",
}

func IsValidChapterStatus(id ChapterStatus) bool {
	switch id {
	case DoneChapter, FindingPersonsChapter, InWorkChapter:
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
	Id              int               `json:"id"`
	MangaName       string            `json:"manga_name"`
	Number          int               `json:"number"`
	Pages           int               `json:"pages"`
	PartsCleanOwner []PagesOwner      `json:"parts_clean_owner"`
	PartsTyperOwner []PagesOwner      `json:"parts_typer_owner"`
	Translator      []PagesOwner      `json:"translator"`
	Editor          []PagesOwner      `json:"editor"`
	Message         discordgo.Message `json:"message"`
	Status          ChapterStatus     `json:"status"`
}

type PagesOwner struct {
	StartPage int
	EndPage   int
	UserId    string
}
