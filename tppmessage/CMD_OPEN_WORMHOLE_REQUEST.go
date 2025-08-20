package tppmessage

type CmdOpenWormholeRequest struct {
	Flag           string `json:"flag"`
	IsOpen         int    `json:"is_open"`
	Msgid          string `json:"msgid"`
	PlayerID       int    `json:"player_id"`
	RetaliateScore int    `json:"retaliate_score"`
	Rqid           int    `json:"rqid"`
	ToPlayerID     int    `json:"to_player_id"`
}
