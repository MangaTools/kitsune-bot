package main

import (
	"github.com/ShaDream/kitsune-bot/group"
	"github.com/ShaDream/kitsune-bot/router"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

var token string

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error occured while reading env file %s", err.Error())
	}
	token = os.Getenv("token")
}

func main() {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		logrus.Fatal("error creating Discord session,", err)
	}

	r := router.NewRouter(s, "!")
	chapterWorker := group.NewChapterWorker()
	chapterWorker.RegisterCommands(r)

	r.Start()

	// In this example, we only care about receiving message events.
	s.Identify.Intents = discordgo.IntentsGuildMessages

	err = s.Open()
	if err != nil {
		logrus.Fatal("can't open connection,", err)
	}
	defer s.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	logrus.Println("Gracefully shutdowning")
}
