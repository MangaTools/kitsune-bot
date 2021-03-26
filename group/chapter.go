package group

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
	Id              int
	MangaName       string
	Number          int
	Pages           int
	Parts           int
	PartsCleanOwner []discordgo.User
	PartsTyperOwner []discordgo.User
	Translator      []discordgo.User
	Editor          []discordgo.User
	Message         discordgo.Message
	Status          ChapterStatus
}
