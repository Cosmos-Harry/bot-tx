package main

import (
	"log"

	"github.com/yanzay/tbot/v2"
)

// func to call the logic and send a message
func (a *application) GenTxHandler(m *tbot.Message) {
	msg, err := GetTokensBulk()
	if err != nil {
		log.Println(err)
	}
	a.client.SendMessage(m.Chat.ID, "Latest 5 spends by AADAO:\n\n"+msg, tbot.OptParseModeHTML)
}
