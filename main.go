package main

import (
	"log"
	"os"
	"time"

	"github.com/yanzay/tbot/v2"
)

type application struct {
	client *tbot.Client
}

var (
	app       application
	bot       *tbot.Server
	tgtoken   string
	newMsg    string
	newTxHash string
)

func main() {

	tgtoken = os.Getenv("TOKEN")
	bot = tbot.New(tgtoken)
	app.client = bot.Client()
	bot.Use(stat)
	bot.HandleMessage("show aadao", app.GenTxHandler)
	go func() {
		for {

			oldMsg, err := GetTokensAuto()
			if err != nil {
				log.Println(err)
			}

			if oldMsg != newMsg && Txhash != newTxHash {
				bot.Use(stat)
				chatID := "@atomgov"
				_, err = app.client.SendMessage(chatID, "Latest spend by AADAO:\n\n"+oldMsg, tbot.OptParseModeHTML)
				if err != nil {
					log.Println("Error sending message:", err)
				}
				newMsg = oldMsg
				newTxHash = Txhash
			}

			time.Sleep(20 * time.Second)
		}
	}()
	log.Fatal(bot.Start())
}
