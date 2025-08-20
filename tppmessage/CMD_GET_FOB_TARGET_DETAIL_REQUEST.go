package tppmessage

type CmdGetFobTargetDetailRequest struct {
	HighRank     int    `json:"high_rank"`
	IsEvent      int    `json:"is_event"`
	IsPlus       int    `json:"is_plus"`
	IsSneak      int    `json:"is_sneak"`
	Mode         string `json:"mode"`
	MotherBaseID int    `json:"mother_base_id"`
	Msgid        string `json:"msgid"`
	Rqid         int    `json:"rqid"`
}
