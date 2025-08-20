package tppmessage

type CmdGetMBCoinRemainderResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Remainder  int         `json:"remainder"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}
