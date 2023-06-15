package main

import (
	"log"
	"os"

	"github.com/yanzay/tbot/v2"
)

type application struct {
	client *tbot.Client
}

var (
	app   application
	bot   *tbot.Server
	tgtoken string
)

func main() {
	// e := godotenv.Load()
	// if e != nil {
	// 	log.Println(e)
	// }

	// token = os.Getenv("TOKEN")

	tgtoken = os.Getenv("TOKEN")
	bot = tbot.New(tgtoken)
	app.client = bot.Client()
	bot.Use(stat)
	bot.HandleMessage("show aadao", app.GenTxHandler)
	log.Fatal(bot.Start())
}
