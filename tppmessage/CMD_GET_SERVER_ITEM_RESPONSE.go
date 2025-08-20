package tppmessage

type CmdGetServerItemResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Item       ServerItem  `json:"item"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}

type ServerItem struct {
	CreateDate int `json:"create_date"`
	Develop    int `json:"develop"`
	Gmp        int `json:"gmp"`
	Id         int `json:"id"`
	LeftSecond int `json:"left_second"`
	MaxSecond  int `json:"max_second"`
	MbCoin     int `json:"mb_coin"`
	Open       int `json:"open"`
}
