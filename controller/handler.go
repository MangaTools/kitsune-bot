package controller

import (
	"github.com/ShaDream/kitsune-bot/models"
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
	r.RegisterMiddleWare(h.SetUserAccess)

	h.RegisterMangaCommands(r)
	h.RegisterChapterCommands(r)
	h.RegisterWorkCommands(r)
	h.RegisterUserCommands(r)
	h.SetUpAccessHandlers(r)

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

func (h *Handler) SetUserAccess(session *discordgo.Session, create *discordgo.MessageCreate, ctx *router.RouterContext, command router.OnMessageCommand) bool {
	if g, err := session.Guild(create.GuildID); err == nil && g.OwnerID == create.Author.ID {
		ctx.UserAccess = models.Admin
		return false
	}
	member, _ := session.GuildMember(create.GuildID, create.Author.ID)
	userRoles := member.Roles
	for _, roleId := range userRoles {
		if roleAccess, ok := AccessRoles[roleId]; ok {
			if roleAccess > ctx.UserAccess {
				ctx.UserAccess = roleAccess
				if ctx.UserAccess == models.Admin {
					return false
				}
			}
		}
	}
	if ctx.UserAccess < command.CommandAccess {
		return true
	}
	return false
}
