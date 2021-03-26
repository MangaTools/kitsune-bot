package router

import "github.com/bwmarrin/discordgo"

type BaseCommand struct {
	Name        string
	Description string
	GroupName   string
	HelpText    string
}

type OnMessageCommand struct {
	BaseCommand
	Handler OnMessageCreate
}

func NewOnMessageCommand(Name string, Description string, GroupName string, HelpText string, handler OnMessageCreate) *OnMessageCommand {
	return &OnMessageCommand{BaseCommand: BaseCommand{
		Name:        Name,
		Description: Description,
		GroupName:   GroupName,
		HelpText:    HelpText},
		Handler: handler}
}

type OnMessageCreate func(*discordgo.Session, *discordgo.MessageCreate, *RouterContext)
