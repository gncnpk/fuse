package tppmessage

type CmdGetRankingRequest struct {
	EventID int    `json:"event_id"`
	GetType string `json:"get_type"`
	Index   int    `json:"index"`
	IsNew   int    `json:"is_new"`
	Msgid   string `json:"msgid"`
	Num     int    `json:"num"`
	Rqid    int    `json:"rqid"`
	Type    string `json:"type"`
}
