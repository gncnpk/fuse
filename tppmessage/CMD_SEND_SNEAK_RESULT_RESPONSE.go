package tppmessage

type SneakResultReward struct {
	AfricanPeach   int   `json:"african_peach"`
	Biotic         int   `json:"biotic"`
	BlackCarrot    int   `json:"black_carrot"`
	Common         int   `json:"common"`
	DigitalisL     int   `json:"digitalis_l"`
	DigitalisP     int   `json:"digitalis_p"`
	Fuel           int   `json:"fuel"`
	Gmp            int   `json:"gmp"`
	GoldenCrescent int   `json:"golden_crescent"`
	Haoma          int   `json:"haoma"`
	IsBefore       int   `json:"is_before"`
	KeyItem        int   `json:"key_item"`
	MainType       int   `json:"main_type"`
	Minor          int   `json:"minor"`
	ParamID        int   `json:"param_id"`
	Precious       int   `json:"precious"`
	Rate           int   `json:"rate"`
	Section        int   `json:"section"`
	StaffCount     int   `json:"staff_count"`
	StaffRank      []int `json:"staff_rank"`
	StaffType      int   `json:"staff_type"`
	Tarragon       int   `json:"tarragon"`
	Wormwood       int   `json:"wormwood"`
}

type CmdSendSneakResultResponse struct {
	CryptoType          string            `json:"crypto_type"`
	EventPoint          int               `json:"event_point"`
	Flowid              any               `json:"flowid"`
	IsSecurityChallenge int               `json:"is_security_challenge"`
	IsWormholeOpen      int               `json:"is_wormhole_open"`
	Msgid               string            `json:"msgid"`
	Result              string            `json:"result"`
	ResultReward        SneakResultReward `json:"result_reward"`
	Rqid                int               `json:"rqid"`
	SneakPoint          int               `json:"sneak_point"`
	Xuid                any               `json:"xuid"`
}
