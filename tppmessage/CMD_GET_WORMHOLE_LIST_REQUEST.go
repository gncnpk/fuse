package tppmessage

type CmdGetWormholeListRequest struct {
	Flag  string `json:"flag"` // always "BLACK"? other possible value is FRIEND or FRIENDLY
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
