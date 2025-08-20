package tppmessage

type CmdUpdateSessionResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	FobIndex   int    `json:"fob_index"`
	Msgid      string `json:"msgid"`
	Result     string `json:"result"`
	Rqid       int    `json:"rqid"`
	SneakMode  int    `json:"sneak_mode"`
	Xuid       any    `json:"xuid"`
}
