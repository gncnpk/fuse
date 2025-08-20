package tppmessage

type CmdAuthSteamticketResponse struct {
	AccountID       string `json:"account_id"`
	CryptoType      string `json:"crypto_type"`
	Currency        string `json:"currency"`
	Flowid          any    `json:"flowid"`
	LoginidPassword string `json:"loginid_password"`
	Msgid           string `json:"msgid"`
	Result          string `json:"result"`
	Rqid            int    `json:"rqid"`
	SmartDeviceID   string `json:"smart_device_id"`
	Xuid            any    `json:"xuid"`
}
