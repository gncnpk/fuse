package tppmessage

type CmdGetResourceParamResponse struct {
	CryptoType          string      `json:"crypto_type"`
	Flowid              interface{} `json:"flowid"`
	Msgid               string      `json:"msgid"`
	NuclearDevelopCosts []int       `json:"nuclear_develop_costs"`
	Result              string      `json:"result"`
	Rqid                int         `json:"rqid"`
	Xuid                interface{} `json:"xuid"`
}
