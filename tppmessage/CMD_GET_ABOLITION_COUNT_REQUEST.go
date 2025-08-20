package tppmessage

type CmdGetAbolitionCountRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
