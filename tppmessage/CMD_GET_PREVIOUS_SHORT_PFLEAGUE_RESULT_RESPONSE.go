package tppmessage

type CmdGetPreviousShortPfleagueResultResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	Info       struct {
		Begin        int   `json:"begin"`
		Current      int   `json:"current"`
		End          int   `json:"end"`
		HistoryCount int   `json:"history_count"`
		MatchHistory []any `json:"match_history"`
		PlayerCount  int   `json:"player_count"`
		PlayerInfo   []any `json:"player_info"`
		Season       int   `json:"season"`
		Section      int   `json:"section"`
	} `json:"info"`
	Msgid  string `json:"msgid"`
	Result string `json:"result"`
	Rqid   int    `json:"rqid"`
	Xuid   any    `json:"xuid"`
}
