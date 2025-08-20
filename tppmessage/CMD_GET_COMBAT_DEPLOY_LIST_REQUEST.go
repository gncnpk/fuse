package tppmessage

type CmdGetCombatDeployListRequest struct {
	Msgid  string `json:"msgid"`
	Num    int    `json:"num"`
	Offset int    `json:"offset"`
	Rqid   int    `json:"rqid"`
}
