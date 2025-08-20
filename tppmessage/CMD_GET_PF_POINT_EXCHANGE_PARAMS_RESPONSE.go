package tppmessage

type CmdGetPfPointExchangeParamsResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	Msgid      string `json:"msgid"`
	ParamList  []struct {
		CommonValue   int `json:"common_value"`
		Count         int `json:"count"`
		ExchangeLimit int `json:"exchange_limit"`
		InfoLangID    int `json:"info_lang_id"`
		LimitedCount  int `json:"limited_count"`
		NameLangID    int `json:"name_lang_id"`
		Point         int `json:"point"`
		Type          int `json:"type"`
		UniqueID      int `json:"unique_id"`
	} `json:"param_list"`
	ParamNum int    `json:"param_num"`
	Result   string `json:"result"`
	Rqid     int    `json:"rqid"`
	Xuid     any    `json:"xuid"`
}
