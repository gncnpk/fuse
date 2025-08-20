package tppmessage

type CmdSyncEmblemResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	Msgid      string `json:"msgid"`
	Result     string `json:"result"`
	Rqid       int    `json:"rqid"`
	Version    int    `json:"version"`
	Xuid       any    `json:"xuid"`
}
