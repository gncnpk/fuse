package tppmessage

type EDefenseMBStatus int

const (
	InfiltratorNoLongerPresent EDefenseMBStatus = iota
	AnotherPlayerIsDefending
	FreeToJoin
)

func GetCmdCheckDefenceMotherbaseResponse() CmdCheckDefenceMotherbaseResponse {
	t := CmdCheckDefenceMotherbaseResponse{}
	t.Msgid = CMD_CHECK_DEFENCE_MOTHERBASE.String()
	t.Result = RESULT_NOERR
	t.CryptoType = CRYPTO_TYPE_COMPOUND

	// TODO ??
	t.CheckResult = FreeToJoin

	return t
}

type CmdCheckDefenceMotherbaseResponse struct {
	CheckResult EDefenseMBStatus `json:"check_result"`
	CryptoType  string           `json:"crypto_type"`
	Flowid      interface{}      `json:"flowid"`
	Msgid       string           `json:"msgid"`
	Result      string           `json:"result"`
	Rqid        int              `json:"rqid"`
	Xuid        interface{}      `json:"xuid"`
}
