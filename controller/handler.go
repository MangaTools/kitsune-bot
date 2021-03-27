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

	h.RegisterMangaCommands(r)
	h.RegisterChapterCommands(r)
	h.RegisterWorkCommands(r)

	r.Start()
}
