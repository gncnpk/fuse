package tppmessage

type CmdSetCurrentplayerResponse struct {
	CryptoType string `json:"crypto_type"`
	Flag       int    `json:"flag"`
	Flowid     any    `json:"flowid"`
	Msgid      string `json:"msgid"`
	PlayerID   int    `json:"player_id"`
	Result     string `json:"result"`
	Rqid       int    `json:"rqid"`
	Xuid       any    `json:"xuid"`
}
