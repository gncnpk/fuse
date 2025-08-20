package tppmessage

type CmdGetPurchasableAreaListResponse struct {
	Area       []CmdGetPurchasableAreaListResponseArea `json:"area"`
	CryptoType string                                  `json:"crypto_type"`
	Flowid     interface{}                             `json:"flowid"`
	Msgid      string                                  `json:"msgid"`
	Result     string                                  `json:"result"`
	Rqid       int                                     `json:"rqid"`
	Xuid       interface{}                             `json:"xuid"`
}

type CmdGetPurchasableAreaListResponseArea struct {
	AreaId         int   `json:"area_id"`
	LocationIndex  int   `json:"location_index"`
	MiningResource []int `json:"mining_resource"`
	Price          int   `json:"price"` // price/2 = mb coin price
}
