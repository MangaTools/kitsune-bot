package controller

import (
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/router"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func (h *Handler) RegisterMangaCommands(r *router.Router) {
	commands := []router.OnMessageCommand{
		{
			BaseCommand: router.BaseCommand{
				Name:        "манга создать",
				Description: "Создаёт новую мангу в списке.",
				GroupName:   "Манга",
				HelpText:    "манга создать <Имя> - добавляет новую мангу в список возможных для перевода.",
			},
			Handler: h.CreateManga,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "манга удалить",
				Description: "Удаляет мангу из списка.",
				GroupName:   "Манга",
				HelpText:    "манга удалить <ID> - удаляет мангу из списка возможных для перевода.",
			},
			Handler: h.DeleteManga,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "манга статусы",
				Description: "позволяет увидеть все доступные для манги статусы",
				GroupName:   "Манга",
				HelpText:    "манга статусы - позволяет увидеть все доступные для манги статусы",
			},
			Handler: h.GetMangaStatuses,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "манга лист",
				Description: "позволяет увидеть несколько манг",
				GroupName:   "Манга",
				HelpText:    "манга лист <страница(опционально, по умолчанию 1)> - позволяет увидеть до 10 манг на странице",
			},
			Handler: h.ListManga,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "манга показать",
				Description: "позволяет увидеть конкретную мангу",
				GroupName:   "Манга",
				HelpText:    "манга показать <ID> - позволяет увидеть конкретную мангу",
			},
			Handler: h.ShowManga,
		},
	}

	r.RegisterOnMessageCommands(commands)
}

type createMangaArgs struct {
	Name string `validate:"min=1"`
}

func (h *Handler) CreateManga(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(createMangaArgs)

	err := FillAndValidateStruct(args, c.Args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	id, err := h.services.MangaMethods.AddManga(args.Name)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Создана манга %s с ID=%d", args.Name, id))
}

type idMangaArgs struct {
	Id int
}

func (h *Handler) DeleteManga(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(idMangaArgs)
	err := CreateFromStringArgs(c.Args, args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	err = h.services.MangaMethods.DeleteManga(args.Id)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Манга с ID %d удалена.", args.Id))
}

func (h *Handler) ShowManga(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(idMangaArgs)
	err := CreateFromStringArgs(c.Args, args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	manga, err := h.services.MangaMethods.GetManga(args.Id)
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
