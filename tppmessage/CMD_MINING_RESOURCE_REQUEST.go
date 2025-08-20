package tppmessage

type CmdMiningResourceRequest struct {
	IsForceMining  int    `json:"is_force_mining"`
	IsForceProcess int    `json:"is_force_process"`
	Msgid          string `json:"msgid"`
	Rqid           int    `json:"rqid"`
}
