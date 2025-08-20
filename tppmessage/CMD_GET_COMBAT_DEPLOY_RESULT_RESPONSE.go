package tppmessage

type CmdGetCombatDeployResultResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	Msgid      string `json:"msgid"`
	Result     string `json:"result"`
	ResultList []any  `json:"result_list"`
	ResultNum  int    `json:"result_num"`
	Rqid       int    `json:"rqid"`
	Xuid       any    `json:"xuid"`
}
