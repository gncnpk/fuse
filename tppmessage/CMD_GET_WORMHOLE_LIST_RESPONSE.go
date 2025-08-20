package tppmessage

// TODO structures guessed from exe, need a real response

type WormholeList struct {
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	Flag       string `json:"flag"` // always "BLACK"? other possible value is FRIEND or FRIENDLY
}

type CmdGetWormholeListResponse struct {
	CryptoType   string         `json:"crypto_type"`
	Flowid       any            `json:"flowid"`
	Msgid        string         `json:"msgid"`
	Result       string         `json:"result"`
	Rqid         int            `json:"rqid"`
	Xuid         any            `json:"xuid"`
	WormholeNum  int            `json:"wormhole_num"`
	WormholeList []WormholeList `json:"wormhole_list"`
}
