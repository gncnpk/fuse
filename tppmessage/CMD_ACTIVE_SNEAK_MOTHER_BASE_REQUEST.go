package tppmessage

type CmdActiveSneakMotherBaseRequest struct {
	MotherBaseID int    `json:"mother_base_id"`
	Msgid        string `json:"msgid"`
	Rqid         int    `json:"rqid"`
}
