package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type application struct {
	client *tbot.Client
}

var (
	app   application
	bot   *tbot.Server
	token string
)

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	token = os.Getenv("TOKEN")
	bot = tbot.New(token)
	app.client = bot.Client()
	bot.Use(stat)
	bot.HandleMessage("", app.GenTxHandler)
	log.Fatal(bot.Start())
}
