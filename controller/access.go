package controller

import (
	"fmt"
	"github.com/ShaDream/kitsune-bot/models"
	"github.com/ShaDream/kitsune-bot/router"
	"github.com/bwmarrin/discordgo"
)

var AccessRoles = map[string]models.RoleAccess{}

func (h *Handler) SetUpAccessHandlers(r *router.Router) {
	commands := []*router.OnMessageCommand{
		{
			BaseCommand: router.BaseCommand{
				Name:          "роли доступа показать",
				Description:   "Показывает роли доступа которые возможно поставить.",
				GroupName:     "Роли доступа",
				HelpText:      "роли доступа показать - показывает роли доступа которые возможно поставить.",
				CommandAccess: models.Admin,
			},
			Handler: h.ShowRoleAccesses,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:          "роли доступа установить",
				Description:   "Даёт определённый доступ роли дискорда.",
				GroupName:     "Роли доступа",
				HelpText:      "роли доступа установить <Id доступа> <Id роли дискорда> - даёт определённый доступ роли дискорда.",
				CommandAccess: models.Admin,
			},
			Handler: h.SetRoleAccess,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:          "роли доступа удалить",
				Description:   "Удаляет определённый доступ у роли дискорда.",
				GroupName:     "Роли доступа",
				HelpText:      "роли доступа удалить <Id роли дискорда> - удаляет определённый доступ у роли дискорда.",
				CommandAccess: models.Admin,
			},
			Handler: h.RemoveRoleAccess,
		},
	}

	r.RegisterOnMessageCommands(commands)
	r.SetGroupAccess("Роли доступа", models.Admin)
	h.SetAccessRoles()
}

func (h *Handler) SetAccessRoles() {
	roles, _ := h.services.AccessMethods.GetAllRoleAccesses()
	for _, role := range roles {
		AccessRoles[role.RoleId] = role.AccessLevel
	}
}

func (h *Handler) ShowRoleAccesses(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Возможные роли доступа:\n%s", models.GetAllRoleAccessesString()))
}

type roleAccessArgs struct {
	Access models.RoleAccess `validate:"min=0,max=4"`
	RoleId string
}

func (h *Handler) SetRoleAccess(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(roleAccessArgs)

	if err := FillAndValidateStruct(args, c.Args); err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	if roles, err := session.GuildRoles(create.GuildID); err == nil {
		hasRole := false
		for _, role := range roles {
			if role.ID == args.RoleId {
				hasRole = true
				break
			}
		}
		if !hasRole {
			session.ChannelMessageSend(create.ChannelID, "Такой роли не существует.")
			return
		}
	}

	if err := h.services.AccessMethods.SetRoleAccess(args.RoleId, args.Access); err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	AccessRoles[args.RoleId] = args.Access
	session.ChannelMessageSend(create.ChannelID, "Роль обновлена/создана.")
}

type roleArgs struct {
	RoleId string
}

func (h *Handler) RemoveRoleAccess(session *discordgo.Session, create *discordgo.MessageCreate, c *router.RouterContext) {
	args := new(roleArgs)

	if err := FillAndValidateStruct(args, c.Args); err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}
	if roles, err := session.GuildRoles(create.GuildID); err == nil {
		hasRole := false
		for _, role := range roles {
			if role.ID == args.RoleId {
				hasRole = true
				break
			}
		}
		if !hasRole {
			session.ChannelMessageSend(create.ChannelID, "Такой роли не существует.")
			return
		}
	}

	if err := h.services.AccessMethods.RemoveRoleAccess(args.RoleId); err != nil {
		session.ChannelMessageSend(create.ChannelID, err.Error())
		return
	}

	delete(AccessRoles, args.RoleId)
	session.ChannelMessageSend(create.ChannelID, "Роль успешно удалена.")
}
