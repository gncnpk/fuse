package tppmessage

type CmdGetFobDamageResponse struct {
	CryptoType string                        `json:"crypto_type"`
	Damage     CmdGetFobDamageResponseDamage `json:"damage"`
	Flowid     interface{}                   `json:"flowid"`
	Msgid      string                        `json:"msgid"`
	Result     string                        `json:"result"`
	Rqid       int                           `json:"rqid"`
	Xuid       interface{}                   `json:"xuid"`
}

type CmdGetFobDamageResponseDamage struct {
	Gmp      int                    `json:"gmp"`
	Resource CmdMiningResourceEntry `json:"resource"`
	Staff    int                    `json:"staff"`
}
