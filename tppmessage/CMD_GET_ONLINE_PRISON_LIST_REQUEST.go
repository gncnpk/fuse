package tppmessage

type CmdGetOnlinePrisonListRequest struct {
	IsPersuade int    `json:"is_persuade"`
	Msgid      string `json:"msgid"`
	Num        int    `json:"num"`
	Offset     int    `json:"offset"`
	Rqid       int    `json:"rqid"`
}
