package controller

import (
	"fmt"
	"github.com/ShaDream/kitsune-bot/router"
	"github.com/bwmarrin/discordgo"
)

func (h *Handler) RegisterChapterCommands(r *router.Router) {
	commands := []router.OnMessageCommand{
		{
			BaseCommand: router.BaseCommand{
				Name:        "глава создать",
				Description: "Создаёт главу для манги.",
				GroupName:   "Глава",
				HelpText:    "глава создать <ID манги> <Номер главы> <Кол-во страниц> - создаёт главу для манги, над которой можно работать и бронировать страницы для работы.",
			},
			Handler: h.CreateChapter,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "глава удалить",
				Description: "Удаляет главу манги",
				GroupName:   "Глава",
				HelpText:    "глава удалить <ID главы> - удаляет главу манги.",
			},
			Handler: h.DeleteChapter,
		},
	}

	r.RegisterOnMessageCommands(commands)
}

type createChapterArgs struct {
	MangaId       int
	ChapterNumber float32 `validate:"min=0"`
	Pages         int     `validate:"min=1"`
}

func (h *Handler) CreateChapter(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(createChapterArgs)
	err := CreateFromStringArgs(c.Args, args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	err = Validator.Struct(args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	chapterId, err := h.services.ChapterMethods.AddChapter(args.MangaId, args.ChapterNumber, args.Pages)

	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Создана новая глава с ID=%d", chapterId))
}

type deleteChapterArgs struct {
	ChapterId int
}

func (h *Handler) DeleteChapter(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(deleteChapterArgs)
	err := CreateFromStringArgs(c.Args, args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	err = h.services.ChapterMethods.DeleteChapter(args.ChapterId)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	session.ChannelMessageSend(create.ChannelID, "Глава успешно удалена!")

}
