package main

import (
	kitsune_bot "github.com/ShaDream/kitsune-bot"
	"github.com/ShaDream/kitsune-bot/controller"
	"github.com/ShaDream/kitsune-bot/repository"
	"github.com/ShaDream/kitsune-bot/service"
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
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error occured while reading env file %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("db_host"),
		Port:     os.Getenv("db_port"),
		Username: os.Getenv("db_user"),
		Password: os.Getenv("db_pass"),
		DBName:   os.Getenv("db_name"),
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("error occured while connecting to db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := controller.NewHandler(services)

	bot := kitsune_bot.NewBot(handlers, token)
	err = bot.Start()
	if err != nil {
		logrus.Fatal(err)
	}
	defer bot.Stop()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	logrus.Println("Gracefully shutdowning...")
}
