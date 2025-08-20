package tppmessage

type CmdAbortMotherBaseRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
