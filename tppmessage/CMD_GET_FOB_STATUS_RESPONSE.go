package tppmessage

type CmdGetFobStatusResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	FobIndex   int    `json:"fob_index"`
	IsRescue   int    `json:"is_rescue"` // emergency mission
	IsReward   int    `json:"is_reward"`
	KillCount  int    `json:"kill_count"`
	Msgid      string `json:"msgid"`
	Record     struct {
		Defense struct {
			Lose int `json:"lose"`
			Win  int `json:"win"`
		} `json:"defense"`
		Insurance  int `json:"insurance"`
		Score      int `json:"score"`
		ShieldDate int `json:"shield_date"`
		Sneak      struct {
			Lose int `json:"lose"`
			Win  int `json:"win"`
		} `json:"sneak"`
	} `json:"record"`
	Result    string `json:"result"`
	Rqid      int    `json:"rqid"`
	SneakMode int    `json:"sneak_mode"`
	Xuid      any    `json:"xuid"`
}
