package tppmessage

type CmdSetCurrentplayerRequest struct {
	Index   int    `json:"index"`
	IsReset int    `json:"is_reset"`
	Msgid   string `json:"msgid"`
	Rqid    int    `json:"rqid"`
}
