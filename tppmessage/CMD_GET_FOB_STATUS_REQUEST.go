package tppmessage

type CmdGetFobStatusRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
