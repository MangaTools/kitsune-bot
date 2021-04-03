package controller

import (
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/router"
	"github.com/bwmarrin/discordgo"
)

func (h *Handler) RegisterUserCommands(r *router.Router) {
	commands := []*router.OnMessageCommand{
		{
			BaseCommand: router.BaseCommand{
				Name:          "юзер статистика",
				Description:   "Показывает статистики пользователя.",
				GroupName:     "Юзер",
				HelpText:      "юзер статистика - показывает статистику пользователя, который вызвал команду.",
				CommandAccess: models.Reader,
			},
			Handler: h.GetUserStatistics,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:          "юзер топ",
				Description:   "Показывает топ пользователей по одной из характеристик.",
				GroupName:     "Юзер",
				HelpText:      "юзер топ <характеристика> - показывает топ пользователей по выбранной характеристике. Чтобы узнать, какие есть характеристики введите команду юзер характеристики",
				CommandAccess: models.Reader,
			},
			Handler: h.GetUserTop,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:          "юзер характеристики",
				Description:   "Показывает все доступные для пользователя характеристики.",
				GroupName:     "Юзер",
				HelpText:      "юзер характеристики - показывает все доступные для пользователя характеристики.",
				CommandAccess: models.Reader,
			},
			Handler: h.GetUserCharacteristics,
		},
	}

	r.RegisterOnMessageCommands(commands)
	r.SetGroupAccess("Юзер", models.Reader)
}

func (h *Handler) GetUserStatistics(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	userId := create.Author.ID
	user, err := h.services.UserMethods.GetUser(userId, create.Author.Username)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	message := fmt.Sprintf("%s\n%s", create.Author.Username, user.GetInfo())
	session.ChannelMessageSend(create.ChannelID, message)

}

type userTopArgs struct {
	Characteristic models.UserCharacteristic `validate:"min=0,max=5"`
}

func (h *Handler) GetUserTop(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(userTopArgs)
	err := FillAndValidateStruct(args, c.Args)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	users, err := h.services.UserMethods.GetTopUser(args.Characteristic)
	if err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	if len(users) == 0 {
		session.ChannelMessageSend(create.ChannelID, "Составить топ пока не возможно.")
		return
	}

	resultText := ""
	for i, user := range users {
		switch args.Characteristic {
		case models.UserCharacteristicScore:
			resultText += fmt.Sprintf("Топ %d: %s - %d\n", i+1, user.Username, user.Score)
		case models.UserCharacteristicTranslatedPages:
			resultText += fmt.Sprintf("Топ %d: %s - %d\n", i+1, user.Username, user.TranslatedPages)
		case models.UserCharacteristicEditedPages:
			resultText += fmt.Sprintf("Топ %d: %s - %d\n", i+1, user.Username, user.EditedPages)
		case models.UserCharacteristicCleanedPages:
			resultText += fmt.Sprintf("Топ %d: %s - %d\n", i+1, user.Username, user.CleanedPages)
		case models.UserCharacteristicTypedPages:
			resultText += fmt.Sprintf("Топ %d: %s - %d\n", i+1, user.Username, user.TypedPages)
		}
	}

	session.ChannelMessageSend(create.ChannelID, resultText)
}

func (h *Handler) GetUserCharacteristics(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	message := fmt.Sprintf("Доступные характеристики:\n%s", models.GetAllUserCharacteristicsString())
	session.ChannelMessageSend(create.ChannelID, message)
}
