package main

import (
	"fmt"

	"github.com/yanzay/tbot/v2"
)

func GetTokens() (string, error) {
	p, err := GetApiCall()
	var res string
	var fin string
	txs := p.TxResponses
	for _, u := range txs {
		logs := u.Logs
		for i, v := range logs {
			events := v.Events
			for _, w := range events {
				if w.Type == "coin_spent" {
					amt := w.Attributes

					for _, x := range amt {

						if x.Key == "amount" {

							res = fmt.Sprintf("%v\n", x.Value)

						}

					}

				}

			}
			fin += res
			i++
		}

	}
	return fin, err
}

func stat(h tbot.UpdateHandler) tbot.UpdateHandler {
	return func(u *tbot.Update) {
		// start := time.Now()
		h(u)
		// log.Printf("Handle time: %v", time.Since(start))
	}
}
