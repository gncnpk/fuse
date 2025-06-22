package tppmessage

type CmdGetFobParamResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	Msgid      string `json:"msgid"`
	Params     []int  `json:"params"`
	Result     string `json:"result"`
	Rqid       int    `json:"rqid"`
	Xuid       any    `json:"xuid"`
}
