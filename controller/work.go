package controller

import (
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/router"
	"github.com/bwmarrin/discordgo"
)

func (h *Handler) RegisterWorkCommands(r *router.Router) {
	commands := []*router.OnMessageCommand{
		{
			BaseCommand: router.BaseCommand{
				Name:          "работа проверить",
				Description:   "Ставит статус забронированных страниц для проверки",
				GroupName:     "Работа",
				HelpText:      "глава проверить <ID работы> - выставляет работу на проверку. Пользуйтесь этой командой, только если вы сделаны страницы, указанные в работе.",
				CommandAccess: models.TeamMember,
			},
			Handler: h.CheckWork,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "работа бронь",
				Description: "Бронирует указанные страницы для работы",
				GroupName:   "Работа",
				HelpText: "глава бронь <ID главы> <Вид работы> <Начальная страница> <Конечная страница> - Бронирует главу для работы." +
					" После брони можете приступать к работе. Для того, чтобы узнать, какой номер у вида работы, выполните команду работа виды",
				CommandAccess: models.TeamMember,
			},
			Handler: h.BookWork,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:          "работа убрать бронь",
				Description:   "Убирает бронь со страниц главы",
				GroupName:     "Работа",
				HelpText:      "главы убрать бронь <ID работы> - убирает бронь с работы. Работает только если работа находиться в статусе \"В процессе\"",
				CommandAccess: models.TeamMember,
			},
			Handler: h.RemoveBookedWork,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:          "работа сделано",
				Description:   "Переводит статус работы в \"Готово\"",
				GroupName:     "Работа",
				HelpText:      "работа сделано <ID работы> - Переводит статус работы в \"Готово\". Работает только с работами, чей статус равен \"На проверке\"",
				CommandAccess: models.Checker,
			},
			Handler: h.DoneWork,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:          "работа виды",
				Description:   "Показывает все доступные виды работ",
				GroupName:     "Работа",
				HelpText:      "работа виды - показывает все доступные виды работ",
				CommandAccess: models.TeamMember,
			},
			Handler: h.GetWorkTypes,
		},
	}

	r.RegisterOnMessageCommands(commands)
	r.SetGroupAccess("Работа", models.TeamMember)
}

type workIdArgs struct {
	ChapterId int
}

func (h *Handler) CheckWork(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(workIdArgs)
	err := FillAndValidateStruct(args, c.Args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	err = h.services.WorkMethods.CheckWork(args.ChapterId)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	session.ChannelMessageSend(create.ChannelID, "Работа выставлена на проверку.")
}

type bookWorkArgs struct {
	ChapterId int
	Type      models.WorkType `validate:"min=0,max=3"`
	StartPage int             `validate:"min=1"`
	EndPage   int             `validate:"gtefield=StartPage"`
}

func (h *Handler) BookWork(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(bookWorkArgs)
	err := FillAndValidateStruct(args, c.Args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	work, err := h.services.WorkMethods.BookWork(create.Author.ID, args.ChapterId, args.Type, args.StartPage, args.EndPage)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Зарезервировано c Id=%d", work))
}

func (h *Handler) RemoveBookedWork(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(workIdArgs)
	err := FillAndValidateStruct(args, c.Args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	err = h.services.WorkMethods.RemoveBookedWork(args.ChapterId)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	session.ChannelMessageSend(create.ChannelID, "Резервирование убрано!")
}

func (h *Handler) DoneWork(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(workIdArgs)
	err := FillAndValidateStruct(args, c.Args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	err = h.services.WorkMethods.DoneWork(args.ChapterId)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	session.ChannelMessageSend(create.ChannelID, "Работа помечена как готовая!")

}

func (h *Handler) GetWorkTypes(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Список доступных работ:\n%s", models.GetAllWorkTypesString()))
}
