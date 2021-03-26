package controller

import (
	"github.com/ShaDream/kitsune-bot/router"
	"github.com/ShaDream/kitsune-bot/service"
	"github.com/bwmarrin/discordgo"
	"os"
)

type Handler struct {
	services *service.Service
	session  *discordgo.Session
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouts(session *discordgo.Session) {
	prefix := os.Getenv("prefix")
	r := router.NewRouter(session, prefix)

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

	for _, c := range commands {
		r.RegisterOnMessageCommand(c)
	}

	r.Start()
}
