package tppmessage

type CmdGetResourceParamRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
