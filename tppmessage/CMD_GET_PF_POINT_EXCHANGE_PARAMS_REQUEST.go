package tppmessage

type CmdGetPfPointExchangeParamsRequest struct {
	IsEvent int    `json:"is_event"`
	Msgid   string `json:"msgid"`
	Rqid    int    `json:"rqid"`
}
