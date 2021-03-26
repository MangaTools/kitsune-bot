package controller

import (
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/router"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func (h *Handler) CreateManga(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	if len(c.Args) < 1 {
		session.ChannelMessageSend(create.ChannelID, "Невозможно создать мангу без имени.")
		return
	}
	name := c.StartText
	id, err := h.services.MangaMethods.AddManga(name)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Создана манга %s с ID=%d", name, id))
}

func (h *Handler) DeleteManga(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	if len(c.Args) < 1 {
		session.ChannelMessageSend(create.ChannelID, "Невозможно удалить мангу без ID.")
		return
	}
	id, err := strconv.Atoi(c.Args[0])
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, "Входные данные - не число.")
		return
	}
	err = h.services.MangaMethods.DeleteManga(id)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Манга с ID %d удалена.", id))
}

func (h *Handler) ShowManga(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	if len(c.Args) < 1 {
		session.ChannelMessageSend(create.ChannelID, "Невозможно показать мангу без ID.")
		return
	}
	id, err := strconv.Atoi(c.Args[0])
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, "Входные данные - не число.")
		return
	}
	manga, err := h.services.MangaMethods.GetManga(id)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	session.ChannelMessageSendEmbed(create.ChannelID, manga.ShowString())
}

func (h *Handler) ListManga(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	var (
		page       = 1
		err  error = nil
	)
	if len(c.Args) > 0 {
		page, err = strconv.Atoi(c.Args[0])
		if err != nil {
			session.ChannelMessageSend(create.ChannelID, "Входные данные - не число.")
			return
		}
		if page < 1 {
			session.ChannelMessageSend(create.ChannelID, "Страница должна быть натуральным числом.")
			return
		}
	}

	mangas, err := h.services.GetMangas(page)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	resutlText := ""
	for _, manga := range mangas {
		resutlText += fmt.Sprintf("%s, id=%d, Статус: %s\n", manga.Name, manga.Id, models.GetMangaStatusString(manga.Status))
	}
	session.ChannelMessageSend(create.ChannelID, resutlText)
}

func (h *Handler) GetMangaStatuses(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	session.ChannelMessageSend(create.ChannelID, models.GetAllMangaStatusesString())
}
