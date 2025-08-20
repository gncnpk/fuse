package tppmessage

type CmdGetOnlinePrisonListResponse struct {
	CryptoType         string `json:"crypto_type"`
	Flowid             any    `json:"flowid"`
	Msgid              string `json:"msgid"`
	PrisonSoldierParam []any  `json:"prison_soldier_param"`
	RescueList         []any  `json:"rescue_list"`
	RescueNum          int    `json:"rescue_num"`
	Result             string `json:"result"`
	Rqid               int    `json:"rqid"`
	SoldierNum         int    `json:"soldier_num"`
	TotalNum           int    `json:"total_num"`
	Xuid               any    `json:"xuid"`
}
