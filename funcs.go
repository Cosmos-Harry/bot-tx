package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/yanzay/tbot/v2"
)

var Txhash string

func GetTokensBulk() (string, error) {

	var res string
	var amt string
	var rcv string
	var fin string
	var pri string
	var new float64

	//Assign both APIs
	t, err := TxApiCall()
	if err != nil {
		fmt.Println(err)
	}
	p, err := PriceApiCall()
	if err != nil {
		fmt.Println(err)
	}

	price := p.Cosmos.Usd

	//The logic to call the amount, value and receiver address from coin_received
	txs := t.TxResponses
	i := 0
	for _, u := range txs {
		if i >= 5 {
			break
		}
		txhash := u.Txhash
		height := u.Height
		timestamp := u.Timestamp.Format("2006-01-02 15:04:05")

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
								rcv = fmt.Sprintf(" to :\n%v", x.Value)
							}
						} else {
							amt = ""
							rcv = ""
						}

					}

				}

			}
		}
		mintScan := fmt.Sprintf("https://www.mintscan.io/cosmos/txs/%v?height=%v", txhash, height)
		shortenedRcv := shortenString(rcv, 26)
		hyperLink := fmt.Sprintf("<a href='%v'>TxLink</a>", mintScan)
		flopri := price * new
		pri = strconv.FormatFloat(flopri, 'f', 2, 64)

		res = "___" + timestamp + "___\n" + amt + "ATOM" + " ($" + pri + ")" + shortenedRcv + "\n" + hyperLink + "\n\n"

		i++
		fin += res

	}
	return fin, err
}
func GetTokensAuto() (string, error) {

	var res string
	var amt string
	var rcv string
	var fin string
	var pri string
	var new float64


	//Assign both APIs
	t, err := TxApiCall()
	if err != nil {
		fmt.Println(err)
	}
	p, err := PriceApiCall()
	if err != nil {
		fmt.Println(err)
	}

	price := p.Cosmos.Usd

	//The logic to call the amount, value and receiver address from coin_received
	txs := t.TxResponses
	i := 0
	for _, u := range txs {
		if i >= 1 {
			break
		}
		Txhash = u.Txhash
		height := u.Height

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
								rcv = fmt.Sprintf(" to :\n%v", x.Value)
							}
						} else {
							amt = ""
							rcv = ""
						}

					}

				}

			}
		}

		mintScan := fmt.Sprintf("https://www.mintscan.io/cosmos/txs/%v?height=%v", Txhash, height)
		shortenedRcv := shortenString(rcv, 26)
		hyperLink := fmt.Sprintf("<a href='%v'>TxLink</a>", mintScan)
		flopri := price * new
		pri = strconv.FormatFloat(flopri, 'f', 2, 64)

		res = amt + "ATOM" + " ($" + pri + ")" + shortenedRcv + "\n" + hyperLink + "\n\n"

		i++
		fin += res

	}
	return fin, err
}

// Logic to update the bot
func stat(h tbot.UpdateHandler) tbot.UpdateHandler {
	return func(u *tbot.Update) {
		h(u)
	}
}

// func to delete characters in a string
func deleteCharacter(s []byte, char byte) []byte {
	return bytes.ReplaceAll(s, []byte{char}, []byte{})
}

// func to return a new string by subtracting characters
func newCharacter(old, chng string) (new string) {

	oldBytes := []byte(old)

	for _, char := range chng {
		oldBytes = deleteCharacter(oldBytes, byte(char))
	}

	newStr := string(oldBytes)

	return newStr

}

func shortenString(s string, maxLength int) string {
	if len(s) > maxLength {
		
		start := s[:15]
		end := s[len(s)-5:]
		return start + "......" + end
	}
	return s
}
