package tppmessage

type CmdSendBootRequest struct {
	IsGoty      int    `json:"is_goty"`
	IsMbcoinDlc int    `json:"is_mbcoin_dlc"`
	Msgid       string `json:"msgid"`
	PlayTime    int    `json:"play_time"`
	Rqid        int    `json:"rqid"`
	SendRecord  struct {
		Config       []int `json:"config"`
		DevelopCount []int `json:"develop_count"`
		Score        []int `json:"score"`
		ScoreLimit   []int `json:"score_limit"`
	} `json:"send_record"`
}
