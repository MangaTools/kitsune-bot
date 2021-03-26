package models

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

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
		return val
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

func (m *Manga) ShowString() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       m.Name,
		Description: fmt.Sprintf("Статус: %s", GetMangaStatusString(m.Status)),
		Color:       0,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "ID",
				Value:  strconv.Itoa(m.Id),
				Inline: false,
			},
		},
	}
	return embed
}

func NewManga(id int, name string) *Manga {
	return &Manga{Name: name, Id: id, Status: TranslatingManga, Chapters: make([]*Chapter, 0)}
}
