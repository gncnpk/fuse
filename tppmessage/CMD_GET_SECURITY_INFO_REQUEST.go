package tppmessage

type CmdGetSecurityInfoRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
