package tppmessage

type CmdGetFobEventListResponse struct {
	CryptoType string                            `json:"crypto_type"`
	EventList  []CmdGetFobEventListResponseEvent `json:"event_list"`
	EventNum   int                               `json:"event_num"`
	Flowid     interface{}                       `json:"flowid"`
	Msgid      string                            `json:"msgid"`
	Result     string                            `json:"result"`
	Rqid       int                               `json:"rqid"`
	Xuid       interface{}                       `json:"xuid"`
}

type CmdGetFobEventListResponseEvent struct {
	AttackerId int `json:"attacker_id"`
	Cluster    int `json:"cluster"`
	EventIndex int `json:"event_index"`
	FobIndex   int `json:"fob_index"`
	IsWin      int `json:"is_win"`
	LayoutCode int `json:"layout_code"`
}
