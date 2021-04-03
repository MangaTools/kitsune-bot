package router

import "github.com/ShaDream/kitsune-bot/models"

type Group struct {
	commands    []*OnMessageCommand
	groupAccess models.RoleAccess
}

func NewGroup() *Group {
	return &Group{commands: make([]*OnMessageCommand, models.Reader)}
}

func (g *Group) AddCommand(command *OnMessageCommand) {
	if command.CommandAccess < g.groupAccess {
		command.CommandAccess = g.groupAccess
	}
	g.commands = append(g.commands, command)
}

func (g *Group) SetGroupAccess(newAccess models.RoleAccess) {
	g.groupAccess = newAccess
	for _, c := range g.commands {
		if c.CommandAccess < g.groupAccess {
			c.CommandAccess = g.groupAccess
		}
	}
}

func (g Group) Commands() (result []OnMessageCommand) {
	result = make([]OnMessageCommand, 0, len(g.commands))
	for _, c := range g.commands {
		result = append(result, *c)
	}
	return
}

func (g Group) GroupAccess() models.RoleAccess {
	return g.groupAccess
}
