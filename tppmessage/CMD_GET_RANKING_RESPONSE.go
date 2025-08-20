package tppmessage

type RankingEntry struct {
	DispRank    int           `json:"disp_rank"`
	FobGrade    int           `json:"fob_grade"`
	IsGradeTop  int           `json:"is_grade_top"`
	LeagueGrade int           `json:"league_grade"`
	PlayerInfo  FobPlayerInfo `json:"player_info"`
	Rank        int           `json:"rank"`
	Score       int           `json:"score"`
}

type CmdGetRankingResponse struct {
	CryptoType  string         `json:"crypto_type"`
	Flowid      any            `json:"flowid"`
	Msgid       string         `json:"msgid"`
	RankingList []RankingEntry `json:"ranking_list"`
	RankingNum  int            `json:"ranking_num"`
	Result      string         `json:"result"`
	Rqid        int            `json:"rqid"`
	UpdateDate  int            `json:"update_date"`
	Xuid        any            `json:"xuid"`
}
