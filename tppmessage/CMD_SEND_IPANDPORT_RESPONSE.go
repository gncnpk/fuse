package tppmessage

type CmdSendIpandportResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	Msgid      string `json:"msgid"`
	Result     string `json:"result"`
	Rqid       int    `json:"rqid"`
	Xuid       any    `json:"xuid"`
}
