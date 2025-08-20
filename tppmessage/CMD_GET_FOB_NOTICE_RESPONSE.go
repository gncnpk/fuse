package tppmessage

type FobNoticeUpdate struct {
	GetPoint  int `json:"get_point"`
	Grade     int `json:"grade"`
	NowRank   int `json:"now_rank"`
	Point     int `json:"point"`
	PrevGrade int `json:"prev_grade"`
	PrevRank  int `json:"prev_rank"`
	Score     int `json:"score"`
}

type FobRecord struct {
	Defense    FobRecordWinLose `json:"defense"`
	Insurance  int              `json:"insurance"`
	Score      int              `json:"score"`
	ShieldDate int              `json:"shield_date"`
	Sneak      FobRecordWinLose `json:"sneak"`
}

type FobRecordWinLose struct {
	Lose int `json:"lose"`
	Win  int `json:"win"`
}

type CmdGetFobNoticeResponse struct {
	ActiveEventServerText        string          `json:"active_event_server_text"`
	CampaignParamList            []any           `json:"campaign_param_list"`
	CommonServerText             string          `json:"common_server_text"`
	CommonServerTextTitle        string          `json:"common_server_text_title"`
	CryptoType                   string          `json:"crypto_type"`
	Daily                        int             `json:"daily"`
	EventDeleteDate              int             `json:"event_delete_date"`
	EventEndDate                 int             `json:"event_end_date"`
	ExistsEventPointCombatDeploy int             `json:"exists_event_point_combat_deploy"`
	Flag                         int             `json:"flag"`
	Flowid                       any             `json:"flowid"`
	LeagueUpdate                 FobNoticeUpdate `json:"league_update"`
	MbCoin                       int             `json:"mb_coin"`
	Msgid                        string          `json:"msgid"`
	PfCurrentSeason              int             `json:"pf_current_season"`
	PfFinishNum                  int             `json:"pf_finish_num"`
	PfFinishNumMax               int             `json:"pf_finish_num_max"`
	PointExchangeEventServerText string          `json:"point_exchange_event_server_text"`
	Record                       FobRecord       `json:"record"`
	Result                       string          `json:"result"`
	Rqid                         int             `json:"rqid"`
	ShortPfCurrentSeason         int             `json:"short_pf_current_season"`
	ShortPfFinishNum             int             `json:"short_pf_finish_num"`
	ShortPfFinishNumMax          int             `json:"short_pf_finish_num_max"`
	SneakUpdate                  FobNoticeUpdate `json:"sneak_update"`
	Xuid                         any             `json:"xuid"`
}
