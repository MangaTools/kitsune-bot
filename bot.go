package kitsune_bot

import (
	"github.com/ShaDream/kitsune-bot/controller"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	handler *controller.Handler
	session *discordgo.Session
	token   string
}

func NewBot(handlers *controller.Handler, token string) *Bot {
	bot := &Bot{handler: handlers, token: token}

	s, err := discordgo.New("Bot " + bot.token)
	if err != nil {
		logrus.Fatal("error creating Discord session,", err)
	}

	bot.session = s
	s.Identify.Intents = discordgo.IntentsGuildMessages

	return bot
}

func (b *Bot) Start() error {
	b.handler.InitRouts(b.session)

	err := b.session.Open()
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) Stop() error {
	err := b.session.Close()
	if err != nil {
		return err
	}
	return nil
}
