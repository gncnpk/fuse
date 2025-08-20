package tppmessage

type CmdGetAbolitionCountResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	Info       struct {
		Count  int `json:"count"`
		Date   int `json:"date"`
		Max    int `json:"max"`
		Num    int `json:"num"`
		Status int `json:"status"`
	} `json:"info"`
	Msgid  string `json:"msgid"`
	Result string `json:"result"`
	Rqid   int    `json:"rqid"`
	Xuid   any    `json:"xuid"`
}
