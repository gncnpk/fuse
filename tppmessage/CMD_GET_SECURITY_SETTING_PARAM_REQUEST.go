package tppmessage

type CmdGetSecuritySettingParamResponse struct {
	CryptoType       string             `json:"crypto_type"`
	Flowid           interface{}        `json:"flowid"`
	Msgid            string             `json:"msgid"`
	Result           string             `json:"result"`
	Rqid             int                `json:"rqid"`
	SecuritySettings []SecuritySettings `json:"security_settings"`
	Xuid             interface{}        `json:"xuid"`
}

type SecuritySettings struct {
	Types []SecuritySettingsTypes `json:"types"`
}

type SecuritySettingsTypes struct {
	LimitNums []int `json:"limit_nums"`
}
