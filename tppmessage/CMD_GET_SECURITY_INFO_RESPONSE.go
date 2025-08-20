package tppmessage

type CmdGetSecurityInfoResponse struct {
	CryptoType      string      `json:"crypto_type"`
	EndDate         int         `json:"end_date"`
	Flowid          interface{} `json:"flowid"`
	InContract      int         `json:"in_contract"`
	InInterval      int         `json:"in_interval"`
	IntervalEndDate int         `json:"interval_end_date"`
	Msgid           string      `json:"msgid"`
	Result          string      `json:"result"`
	Rqid            int         `json:"rqid"`
	Xuid            interface{} `json:"xuid"`
}
