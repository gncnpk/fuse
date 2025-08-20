package tppmessage

type CMDReqAuthHTTPSRequest struct {
	Hash     string `json:"hash"`
	IsTpp    int    `json:"is_tpp"`
	Msgid    string `json:"msgid"`
	Platform string `json:"platform"`
	Rqid     int    `json:"rqid"`
	Ugc      int    `json:"ugc"`
	UserName string `json:"user_name"` // steamID
	Ver      string `json:"ver"`
}

type CMDReqAuthHTTPSResponse struct {
	AesKey        interface{} `json:"aes_key"`
	CbcIv         interface{} `json:"cbc_iv"`
	CryptoKey     string      `json:"crypto_key"`
	CryptoType    string      `json:"crypto_type"`
	Flowid        interface{} `json:"flowid"`
	HeartbeatSec  int         `json:"heartbeat_sec"`
	HmacKey       interface{} `json:"hmac_key"`
	InquiryId     int         `json:"inquiry_id"`
	IsUseApr      int         `json:"is_use_apr"`
	Msgid         string      `json:"msgid"`
	Result        string      `json:"result"`
	Rqid          int         `json:"rqid"`
	Session       string      `json:"session"`
	SmartDeviceID string      `json:"smart_device_id"`
	TimeoutSec    int         `json:"timeout_sec"`
	UserID        int         `json:"user_id"` // user id on server, NOT steamID
	Xuid          interface{} `json:"xuid"`
}
