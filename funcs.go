package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/yanzay/tbot/v2"
)

func GetTokens() (string, error) {

	var res string
	var amt string
	var rcv string
	var fin string
	var pri string
	var new float64

	t, err := TxApiCall()
	if err != nil {
		fmt.Println(err)
	}
	p, err := PriceApiCall()
	if err != nil {
		fmt.Println(err)
	}

	price := p.Cosmos.Usd
	txs := t.TxResponses

	i := 0
	for _, u := range txs {
		if i >= 5 {
			break
		}
		logs := u.Logs
		for _, v := range logs {

			events := v.Events
			for _, w := range events {

				amount := w.Attributes
				if w.Type == "coin_received" {

					for _, x := range amount {

						if x.Value >= "0" {
							if x.Key == "amount" {

								str := newCharacter(x.Value, "uatom")
								in, err := strconv.ParseFloat(str, 64)
								if err != nil {
									fmt.Println(err)
								}
								new = (in / 1000000)
								amt = fmt.Sprintf("%v", new)
							}

							if x.Key == "receiver" {
								rcv = fmt.Sprintf(" to :\t%v\n\n", x.Value)
							}
						} else {
							amt = ""
							rcv = ""
						}

					}

				}

			}
			flopri := price * new
			pri = strconv.FormatFloat(flopri, 'f', 2, 64)

			res = amt + "ATOM" + " ($" + pri + ")" + rcv

			i++
			fin += res
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

func deleteCharacter(s []byte, char byte) []byte {
	return bytes.ReplaceAll(s, []byte{char}, []byte{})
}

func newCharacter(old, chng string) (new string) {

	// Convert string to byte slice
	oldBytes := []byte(old)

	// Delete characters from the byte slice
	for _, char := range chng {
		oldBytes = deleteCharacter(oldBytes, byte(char))
	}

	// Convert byte slice back to string
	newStr := string(oldBytes)

	return newStr

}
