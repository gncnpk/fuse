package tppmessage

type CmdCheckShortPfleagueEnterableResponse struct {
	CryptoType        string `json:"crypto_type"`
	Flowid            any    `json:"flowid"`
	Msgid             string `json:"msgid"`
	PfleagueDate      int    `json:"pfleague_date"`
	Result            string `json:"result"`
	ResultAlreadyRead int    `json:"result_already_read"`
	Rqid              int    `json:"rqid"`
	Status            string `json:"status"`
	Xuid              any    `json:"xuid"`
}
