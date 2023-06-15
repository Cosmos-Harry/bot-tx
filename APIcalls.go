package main

import (
	"encoding/json"
	"net/http"
)

func TxApiCall() (*Tx, error) {
	resp, err := http.Get("https://api-cosmoshub.pupmos.network/cosmos/tx/v1beta1/txs?events=message.sender%3D%27cosmos1jggracsvp6fkw0ktf544lpda6dwrrxggx9a20n%27&order_by=ORDER_BY_DESC")
	p := &Tx{}

	if err != nil {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return p, err
}
func PriceApiCall() (*Price, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=cosmos&vs_currencies=usd")
	p := &Price{}

	if err != nil {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return p, err
}
