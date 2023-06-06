package main

import "time"

type Model struct {
	Txs []struct {
		Body struct {
			Messages []struct {
				Type             string `json:"@type"`
				DelegatorAddress string `json:"delegator_address"`
				ValidatorAddress string `json:"validator_address"`
			} `json:"messages"`
			Memo                        string        `json:"memo"`
			TimeoutHeight               string        `json:"timeout_height"`
			ExtensionOptions            []interface{} `json:"extension_options"`
			NonCriticalExtensionOptions []interface{} `json:"non_critical_extension_options"`
		} `json:"body"`
		AuthInfo struct {
			SignerInfos []struct {
				PublicKey struct {
					Type string `json:"@type"`
					Key  string `json:"key"`
				} `json:"public_key"`
				ModeInfo struct {
					Single struct {
						Mode string `json:"mode"`
					} `json:"single"`
				} `json:"mode_info"`
				Sequence string `json:"sequence"`
			} `json:"signer_infos"`
			Fee struct {
				Amount []struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"amount"`
				GasLimit string `json:"gas_limit"`
				Payer    string `json:"payer"`
				Granter  string `json:"granter"`
			} `json:"fee"`
		} `json:"auth_info"`
		Signatures []string `json:"signatures"`
	} `json:"txs"`
	TxResponses []struct {
		Height    string `json:"height"`
		Txhash    string `json:"txhash"`
		Codespace string `json:"codespace"`
		Code      int    `json:"code"`
		Data      string `json:"data"`
		RawLog    string `json:"raw_log"`
		Logs      []struct {
			MsgIndex int    `json:"msg_index"`
			Log      string `json:"log"`
			Events   []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"events"`
		} `json:"logs"`
		Info      string `json:"info"`
		GasWanted string `json:"gas_wanted"`
		GasUsed   string `json:"gas_used"`
		Tx        struct {
			Type string `json:"@type"`
			Body struct {
				Messages []struct {
					Type             string `json:"@type"`
					DelegatorAddress string `json:"delegator_address"`
					ValidatorAddress string `json:"validator_address"`
				} `json:"messages"`
				Memo                        string        `json:"memo"`
				TimeoutHeight               string        `json:"timeout_height"`
				ExtensionOptions            []interface{} `json:"extension_options"`
				NonCriticalExtensionOptions []interface{} `json:"non_critical_extension_options"`
			} `json:"body"`
			AuthInfo struct {
				SignerInfos []struct {
					PublicKey struct {
						Type string `json:"@type"`
						Key  string `json:"key"`
					} `json:"public_key"`
					ModeInfo struct {
						Single struct {
							Mode string `json:"mode"`
						} `json:"single"`
					} `json:"mode_info"`
					Sequence string `json:"sequence"`
				} `json:"signer_infos"`
				Fee struct {
					Amount []struct {
						Denom  string `json:"denom"`
						Amount string `json:"amount"`
					} `json:"amount"`
					GasLimit string `json:"gas_limit"`
					Payer    string `json:"payer"`
					Granter  string `json:"granter"`
				} `json:"fee"`
			} `json:"auth_info"`
			Signatures []string `json:"signatures"`
		} `json:"tx"`
		Timestamp time.Time `json:"timestamp"`
		Events    []struct {
			Type       string `json:"type"`
			Attributes []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
				Index bool   `json:"index"`
			} `json:"attributes"`
		} `json:"events"`
	} `json:"tx_responses"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}
