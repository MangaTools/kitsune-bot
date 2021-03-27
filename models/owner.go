package models

import "fmt"

type OwnerPageStatus int

type WorkType int

type Owner struct {
	Id        int             `json:"id" db:"id"`
	UserId    string          `json:"user_id" db:"user_id"`
	PageStart int             `json:"page_start" db:"page_start"`
	PageEnd   int             `json:"page_end" db:"page_end"`
	Type      WorkType        `json:"work_type" db:"work_type"`
	Status    OwnerPageStatus `json:"status" db:"status"`
}

const (
	Clean = iota
	Edit
	Type
	Translate
)

var workTypeToString = map[WorkType]string{
	Clean:     "Клин",
	Edit:      "Редакт",
	Type:      "Тайп",
	Translate: "Перевод",
}

const (
	Done = iota
	InProgress
	OnCheck
	OnCompletion
)

var ownerPageStatusToString = map[OwnerPageStatus]string{
	Done:         "Готово",
	InProgress:   "В процессе",
	OnCheck:      "На проверке",
	OnCompletion: "На доработке",
}

func IsValidOwnerPageStatus(id OwnerPageStatus) bool {
	switch id {
	case Done, InProgress, OnCheck, OnCompletion:
		return true
	}
	return false
}

func GetOwnerPageStatusString(id OwnerPageStatus) string {
	if val, ok := ownerPageStatusToString[id]; ok {
		return fmt.Sprintf("%d - %s", int(id), val)
	}
	return "Такого статуса нет"
}

func GetAllOwnerPageStatusesString() string {
	result := ""
	for id, val := range ownerPageStatusToString {
		result += fmt.Sprintf("%d - %s\n", id, val)
	}
	return result
}

func IsValidWorkType(id WorkType) bool {
	switch id {
	case Clean, Edit, Type, Translate:
		return true
	}
	return false
}

func GetWorkTypeString(id WorkType) string {
	if val, ok := workTypeToString[id]; ok {
		return fmt.Sprintf("%d - %s", int(id), val)
	}
	return "Такого статуса нет"
}

func GetAllWorkTypesString() string {
	result := ""
	for id, val := range workTypeToString {
		result += fmt.Sprintf("%d - %s\n", id, val)
	}
	return result
}
