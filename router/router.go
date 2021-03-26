package router

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"strings"
)

type Router struct {
	session *discordgo.Session
	Prefix  string
	routs   map[string]OnMessageCommand
	groups  map[string][]OnMessageCommand
}

func NewRouter(session *discordgo.Session, prefix string) *Router {
	r := &Router{
		Prefix:  prefix,
		session: session,
		routs:   make(map[string]OnMessageCommand),
		groups:  make(map[string][]OnMessageCommand),
	}
	r.setHelpCommand()
	return r
}

func (r *Router) Start() {
	r.session.AddHandler(r.handleMessage)
}

func (r *Router) RegisterOnMessageCommand(command OnMessageCommand) {
	name := command.Name
	if _, ok := r.groups[command.GroupName]; !ok {
		r.groups[command.GroupName] = make([]OnMessageCommand, 0)
	}
	r.groups[command.GroupName] = append(r.groups[command.GroupName], command)

	r.routs[name] = command
}

func (r *Router) setHelpCommand() {
	command := OnMessageCommand{
		BaseCommand: BaseCommand{
			Name:        "помощь",
			Description: fmt.Sprintf("Показывает вам текущую подсказку. Для подробной информации по какой-либо команде напишите %sпомощь \"название команды\"", r.Prefix),
			HelpText:    fmt.Sprintf("Зачем тебе это? Просто вызови %sпомощь.", r.Prefix),
			GroupName:   "",
		},
		Handler: func(session *discordgo.Session, create *discordgo.MessageCreate, arg string) {
			arg = strings.TrimSpace(arg)

			if len(arg) > 0 {
				lowered := strings.ToLower(arg)
				for name, c := range r.routs {
					if strings.HasPrefix(lowered, name) {
						_, err := session.ChannelMessageSend(create.ChannelID, c.HelpText)
						if err != nil {
							logrus.Error(err)
						}
						return
					}
				}
			}

			text := ""
			for groupName, g := range r.groups {
				if groupName != "" {
					text += fmt.Sprintf("%s:\n", groupName)
				}
				for _, c := range g {
					text += fmt.Sprintf("%s%s - %s\n", r.Prefix, c.Name, c.Description)
				}
				text += "\n"
			}

			_, err := session.ChannelMessageSend(create.ChannelID, text)
			if err != nil {
				logrus.Error(err)
			}
		},
	}
	r.RegisterOnMessageCommand(command)
}

func (r *Router) handleMessage(session *discordgo.Session, create *discordgo.MessageCreate) {
	if create.Author.ID == session.State.User.ID {
		return
	}
	logrus.WithFields(logrus.Fields{
		"Username": create.Author.Username,
		"Id":       create.Author.ID,
		"Roles":    create.Member.Roles,
	}).Info("message sent")
	text := create.Content
	if !strings.HasPrefix(text, r.Prefix) {
		return
	}
	text = strings.TrimLeft(text, r.Prefix)
	lowerText := strings.ToLower(text)
	for name, command := range r.routs {
		if strings.HasPrefix(lowerText, name) {
			text = strings.TrimSpace(text[len(name):])
			command.Handler(session, create, text)
			return
		}
	}
}
