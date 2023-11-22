package models

const CoinremitterEndpoint = "https://coinremitter.com/api/v3/BTC/validate-address"

type CoinremitterRequest struct {
	ApiKey   string `json:"api_key"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type CoinremitterResponse struct {
	Flag   int    `json:"flag"`
	Msg    string `json:"msg"`
	Action string `json:"action"`
	Data   struct {
		Valid bool `json:"valid"`
	} `json:"data"`
}
