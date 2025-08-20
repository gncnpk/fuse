package tppmessage

type CmdGetCombatDeployResultRequest struct {
	Msgid   string `json:"msgid"`
	Rqid    int    `json:"rqid"`
	Version int    `json:"version"`
}
