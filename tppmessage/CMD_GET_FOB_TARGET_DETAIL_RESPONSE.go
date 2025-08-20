package tppmessage

type FobWeaponPlacement struct {
	EmplacementGunEast int `json:"emplacement_gun_east"`
	EmplacementGunWest int `json:"emplacement_gun_west"`
	GatlingGun         int `json:"gatling_gun"`
	GatlingGunEast     int `json:"gatling_gun_east"`
	GatlingGunWest     int `json:"gatling_gun_west"`
	MortarNormal       int `json:"mortar_normal"`
}

type FobRewardInfo struct {
	BottomType int `json:"bottom_type"`
	Rate       int `json:"rate"`
	Section    int `json:"section"`
	Type       int `json:"type"`
	Value      int `json:"value"`
}

type FobPrimaryReward struct {
	RewardInfo []FobRewardInfo `json:"reward_info"`
}

type FobSectionStaff struct {
	Base     int `json:"base"`
	Combat   int `json:"combat"`
	Develop  int `json:"develop"`
	Medical  int `json:"medical"`
	Security int `json:"security"`
	Spy      int `json:"spy"`
	Suport   int `json:"suport"`
}

type FobDetail struct {
	CapturedRankBottom  int                `json:"captured_rank_bottom"`
	CapturedRankTop     int                `json:"captured_rank_top"`
	CapturedStaff       int                `json:"captured_staff"`
	MotherBaseParam     MotherBaseParam    `json:"mother_base_param"`
	OwnerPlayerID       int                `json:"owner_player_id"`
	Placement           FobWeaponPlacement `json:"placement"`
	Platform            int                `json:"platform"`
	PrimaryReward       []FobPrimaryReward `json:"primary_reward"`
	RewardRate          int                `json:"reward_rate"`
	SectionStaff        []FobSectionStaff  `json:"section_staff"`
	SecuritySectionRank int                `json:"security_section_rank"`
}

type FobSession struct {
	Ip                  string `json:"ip"`
	IsInvalid           int    `json:"is_invalid"`
	Npid                Npid   `json:"npid"`
	Port                int    `json:"port"`
	SecureDeviceAddress string `json:"secure_device_address"`
	Steamid             uint64 `json:"steamid"`
	Xnaddr              string `json:"xnaddr"`
	Xnkey               any    `json:"xnkey"`
	Xnkid               any    `json:"xnkid"`
	Xuid                uint64 `json:"xuid"`
}

type CmdGetFobTargetDetailResponse struct {
	CryptoType    string     `json:"crypto_type"`
	Detail        FobDetail  `json:"detail"`
	EventClearBit int        `json:"event_clear_bit"`
	Flowid        any        `json:"flowid"`
	IsRestrict    int        `json:"is_restrict"`
	Msgid         string     `json:"msgid"`
	Result        string     `json:"result"`
	Rqid          int        `json:"rqid"`
	Session       FobSession `json:"session"`
	Xuid          any        `json:"xuid"`
}
