package models

import "fmt"

type RoleAccess int

const (
	Reader RoleAccess = iota
	TeamMember
	Checker
	Moderator
	Admin
)

var roleAccessToString = map[RoleAccess]string{
	Reader:     "Читатель",
	TeamMember: "Участник команды",
	Checker:    "Проверяющий",
	Moderator:  "Модератор",
	Admin:      "Админ",
}

type Role struct {
	Id          int
	RoleId      string
	AccessLevel RoleAccess
}

func IsValidRoleAccess(id RoleAccess) bool {
	switch id {
	case Reader, TeamMember, Checker, Moderator, Admin:
		return true
	}
	return false
}

func GetRoleAccessString(id RoleAccess) string {
	if val, ok := roleAccessToString[id]; ok {
		return fmt.Sprintf("%d - %s", int(id), val)
	}
	return "Такого статуса нет"
}

func GetAllRoleAccessesString() string {
	result := ""
	for id, val := range roleAccessToString {
		result += fmt.Sprintf("%d - %s\n", id, val)
	}
	return result
}
