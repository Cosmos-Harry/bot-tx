package main

import (
	"encoding/json"
	"net/http"
)

func GetApiCall() (*Model, error) {
	resp, err := http.Get("https://api-juno.pupmos.network/cosmos/tx/v1beta1/txs?events=message.sender%3D%27juno1lk0cc3ygg9evhlcx4j63p7q46xht03j93e3ug0%27&order_by=ORDER_BY_UNSPECIFIED")
	p := &Model{}

	if err != nil {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return p, err
}
