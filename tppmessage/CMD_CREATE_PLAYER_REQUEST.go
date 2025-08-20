package tppmessage

type CmdCreatePlayerRequest struct {
	Msgid      string `json:"msgid"`
	PlayerName string `json:"player_name"`
	Rqid       int    `json:"rqid"`
}
