package tppmessage

type CmdCheckDefenceMotherbaseRequest struct {
	Msgid         string `json:"msgid"`
	OwnerPlayerId int    `json:"owner_player_id"`
	Rqid          int    `json:"rqid"`
}
