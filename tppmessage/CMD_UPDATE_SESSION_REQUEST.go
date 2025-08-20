package tppmessage

type CmdUpdateSessionRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
