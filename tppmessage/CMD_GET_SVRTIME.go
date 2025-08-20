package tppmessage

type CmdGetSvrTimeRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}

type CmdGetSvrTimeResponse struct {
	CryptoType string      `json:"crypto_type"`
	Date       int         `json:"date"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}
