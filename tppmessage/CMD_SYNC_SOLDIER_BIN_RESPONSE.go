package tppmessage

type CmdSyncSoldierBinResponse struct {
	CryptoType   string `json:"crypto_type"`
	Flowid       any    `json:"flowid"`
	Msgid        string `json:"msgid"`
	Result       string `json:"result"`
	Rqid         int    `json:"rqid"`
	SoldierNum   int    `json:"soldier_num"` // 1 soldier = 16 bytes
	SoldierParam string `json:"soldier_param"`
	Version      int    `json:"version"`
	Xuid         any    `json:"xuid"`
}
