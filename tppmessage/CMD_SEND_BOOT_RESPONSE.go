package tppmessage

type CmdSendBootResponse struct {
	CryptoType string `json:"crypto_type"`
	Flag       int    `json:"flag"`
	Flowid     any    `json:"flowid"`
	Msgid      string `json:"msgid"`
	Result     string `json:"result"`
	Rqid       int    `json:"rqid"`
	Xuid       any    `json:"xuid"`
}
