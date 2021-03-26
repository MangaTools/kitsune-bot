package group

import (
	"encoding/json"
	"fmt"
	"github.com/ShaDream/kitsune-bot/router"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
	"strings"
)

type MangaWorker struct {
	mangas map[int]*Manga
}

func NewChapterWorker() *MangaWorker {
	return &MangaWorker{mangas: make(map[int]*Manga, 0)}
}

func (m *MangaWorker) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var worker MangaWorker
	err = json.Unmarshal(data, &worker)
	if err != nil {
		return err
	}
	return nil
}

func (m *MangaWorker) RegisterCommands(r *router.Router) {
	commands := []router.OnMessageCommand{
		{
			BaseCommand: router.BaseCommand{
				Name:        "манга создать",
				Description: "Создаёт новую мангу в списке.",
				GroupName:   "Манга",
				HelpText:    "манга создать <Имя> - добавляет новую мангу в список возможных для перевода.",
			},
			Handler: m.CreateManga,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "манга удалить",
				Description: "Удаляет мангу из списка.",
				GroupName:   "Манга",
				HelpText:    "манга удалить <ID> - удаляет мангу из списка возможных для перевода.",
			},
			Handler: m.DeleteManga,
		},
		{
			BaseCommand: router.BaseCommand{
				Name:        "манга статусы",
				Description: "позволяет увидеть все доступные для манги статусы",
				GroupName:   "Манга",
				HelpText:    "манга статусы - позволяет увидеть все доступные для манги статусы",
			},
			Handler: m.GetMangaStatuses,
		},
	}

	for _, c := range commands {
		r.RegisterOnMessageCommand(c)
	}
}

func (m *MangaWorker) GetMangaStatuses(session *discordgo.Session, create *discordgo.MessageCreate, _ string) {
	session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Доступные статусы для манги:\n%s", GetAllMangaStatusesString()))
}

func (m *MangaWorker) DeleteManga(session *discordgo.Session, create *discordgo.MessageCreate, arg string) {
	args := strings.Split(arg, " ")
	resultedArgs := make([]string, 0)
	for _, ar := range args {
		if len(ar) != 0 {
			resultedArgs = append(resultedArgs, ar)
		}
	}

	if len(resultedArgs) == 0 {
		return
	}

	id, err := strconv.Atoi(resultedArgs[0])
	if err != nil {
		return
	}
	if _, ok := m.mangas[id]; !ok {
		_, err = session.ChannelMessageSend(create.ChannelID, "Такой манги не существует!")
		if err != nil {
			logrus.Error(err)
		}
		return
	}
	_, err = session.ChannelMessageSend(create.ChannelID, "Манга успешно удалена!")
	if err != nil {
		logrus.Error(err)
	}
	delete(m.mangas, id)
}

func (m *MangaWorker) CreateManga(session *discordgo.Session, create *discordgo.MessageCreate, arg string) {
	args := strings.Split(arg, " ")
	resultedArgs := make([]string, 0)
	for _, ar := range args {
		if len(ar) != 0 {
			resultedArgs = append(resultedArgs, ar)
		}
	}

	if len(resultedArgs) < 1 {
		return
	}
	name := strings.Join(resultedArgs, " ")
	id := m.CreateId()
	m.mangas[id] = NewManga(id, name)

	_, err := session.ChannelMessageSend(create.ChannelID, fmt.Sprintf("Манга создана с ID = %d", id))
	if err != nil {
		logrus.Error(err)
	}
}

func (m *MangaWorker) CreateId() int {
	id := 0
	for {
		if _, ok := m.mangas[id]; !ok {
			return id
		}
		id++
	}
}
