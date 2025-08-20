package tppmessage

type ServerItemListEntry struct {
	CreateDate int `json:"create_date"`
	Develop    int `json:"develop"`
	Gmp        int `json:"gmp"`
	ID         int `json:"id"`
	LeftSecond int `json:"left_second"`
	MaxSecond  int `json:"max_second"`
	MbCoin     int `json:"mb_coin"`
	Open       int `json:"open"`
}

type CmdGetServerItemListResponse struct {
	CryptoType   string                `json:"crypto_type"`
	DevelopLimit int                   `json:"develop_limit"`
	Flowid       any                   `json:"flowid"`
	ItemList     []ServerItemListEntry `json:"item_list"`
	ItemNum      int                   `json:"item_num"`
	Msgid        string                `json:"msgid"`
	Result       string                `json:"result"`
	Rqid         int                   `json:"rqid"`
	Xuid         any                   `json:"xuid"`
}
