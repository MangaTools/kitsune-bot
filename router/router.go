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

func (r *Router) RegisterOnMessageCommands(commands []OnMessageCommand) {
	for _, c := range commands {
		r.RegisterOnMessageCommand(c)
	}
}

func (r *Router) setHelpCommand() {
	command := OnMessageCommand{
		BaseCommand: BaseCommand{
			Name:        "помощь",
			Description: fmt.Sprintf("Показывает вам текущую подсказку. Для подробной информации по какой-либо команде напишите %sпомощь \"название команды\"", r.Prefix),
			HelpText:    fmt.Sprintf("Зачем тебе это? Просто вызови %sпомощь.", r.Prefix),
			GroupName:   "",
		},
		Handler: r.helpFunction,
	}
	r.RegisterOnMessageCommand(command)
}

func (r *Router) handleMessage(session *discordgo.Session, create *discordgo.MessageCreate) {
	if create.Author.ID == session.State.User.ID {
		return
	}

	// Find prefix and delete
	text := create.Content
	if !strings.HasPrefix(text, r.Prefix) {
		return
	}
	text = strings.TrimLeft(text, r.Prefix)
	lowerText := strings.ToLower(text)

	// Try find command and execute it
	for name, command := range r.routs {
		if strings.HasPrefix(lowerText, name) {
			// remove command text from text content to parse the rest of the line
			text = strings.TrimSpace(text[len(name):])
			command.Handler(session, create, NewRouterContext(text))
			return
		}
	}
}

// built-in function to print help string in chat
func (r *Router) helpFunction(session *discordgo.Session, create *discordgo.MessageCreate, c *RouterContext) {
	if len(c.StartText) > 0 {
		lowered := strings.ToLower(c.StartText)
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
}
