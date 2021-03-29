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

	r.RegisterMiddleWare(h.CreateUserIfDoesntExistsMiddleWare)
	h.RegisterMangaCommands(r)
	h.RegisterChapterCommands(r)
	h.RegisterWorkCommands(r)
	h.RegisterUserCommands(r)

	r.Start()
}

func (h *Handler) CreateUserIfDoesntExistsMiddleWare(session *discordgo.Session, create *discordgo.MessageCreate, ctx *router.RouterContext, command router.OnMessageCommand) bool {
	err := h.services.UserMethods.TryCreateUser(create.Author.ID, create.Author.Username)

	if err != nil {
		session.ChannelMessageSend(create.ChannelID, "Произошла внутренняя ошибка бота.")
		return true
	}
	return false
}
