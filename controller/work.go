package controller

import "github.com/ShaDream/kitsune-bot/router"

func (h *Handler) RegisterWorkCommands(r *router.Router) {
	commands := []router.OnMessageCommand{
		{
			BaseCommand: router.BaseCommand{
				Name:        "работа проверить",
				Description: "Ставит статус забронированных страниц для проверки",
				GroupName:   "Работа",
				HelpText:    "глава проверить <ID работы> - выставляет работу на проверку. Пользуйтесь этой командой, только если вы сделаны страницы, указанные в работе.",
			},
			Handler: nil,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "работа бронь",
				Description: "Бронирует указанные страницы для работы",
				GroupName:   "Работа",
				HelpText: "глава бронь <ID главы> <Вид работы> <Начальная страница> <Конечная страница> - Бронирует главу для работы." +
					" После брони можете приступать к работе. Для того, чтобы узнать, какой номер у вида работы, выполните команду глава виды работы",
			},
			Handler: nil,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "работа убрать бронь",
				Description: "Убирает бронь со страниц главы",
				GroupName:   "Работа",
				HelpText:    "главы убрать бронь <ID работы> - убирает бронь с работы. Работает только если работа находиться в статусе \"В процессе\"",
			},
			Handler: nil,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "работа сделано",
				Description: "Переводит статус работы в \"Готово\"",
				GroupName:   "Работа",
				HelpText:    "работа сделано <ID работы> - Переводит статус работы в \"Готово\". Работает только с работами, чей статус равен \"На проверке\"",
			},
			Handler: nil,
		},
	}

	r.RegisterOnMessageCommands(commands)
}
