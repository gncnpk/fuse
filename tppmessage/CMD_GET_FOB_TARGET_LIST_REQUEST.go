package tppmessage

type CmdGetFobTargetListRequest struct {
	Index int    `json:"index"`
	Msgid string `json:"msgid"`
	Num   int    `json:"num"`
	Rqid  int    `json:"rqid"`
	Type  string `json:"type"`
}
