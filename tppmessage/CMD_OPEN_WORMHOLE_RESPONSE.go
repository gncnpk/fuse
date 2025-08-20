package tppmessage

type CmdOpenWormholeResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	IsNewOpen  int    `json:"is_new_open"`
	Msgid      string `json:"msgid"`
	PlayerID   int    `json:"player_id"`
	Result     string `json:"result"`
	Rqid       int    `json:"rqid"`
	ToPlayerID int    `json:"to_player_id"`
	Xuid       any    `json:"xuid"`
}
