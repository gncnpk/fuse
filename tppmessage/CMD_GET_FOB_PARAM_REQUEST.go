package tppmessage

type CmdGetFobParamRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
