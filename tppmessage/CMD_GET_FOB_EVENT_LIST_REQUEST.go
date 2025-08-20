package tppmessage

type CmdGetFobEventListRequest struct {
	Msgid   string `json:"msgid"`
	Rqid    int    `json:"rqid"`
	Version int    `json:"version"`
}
