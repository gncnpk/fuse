package tppmessage

type CmdGetSvrListResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	ServerNum  int         `json:"server_num"`
	Svrlist    []string    `json:"svrlist"`
	Xuid       interface{} `json:"xuid"`
}

type CmdGetSrvListRequest struct {
	Lang  string `json:"lang"`
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
