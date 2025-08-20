package tppmessage

type CmdGetFobNoticeRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
