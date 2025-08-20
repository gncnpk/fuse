package tppmessage

type FobDeployDamageParam struct {
	ClusterIndex   int   `json:"cluster_index"`
	DamageValues   []int `json:"damage_values"`
	ExpirationDate int   `json:"expiration_date"`
	MotherbaseID   int   `json:"motherbase_id"`
}

type Espionage struct {
	Lose    int `json:"lose"`
	Score   int `json:"score"`
	Section int `json:"section"`
	Win     int `json:"win"`
}

type NpidHandler struct {
	Data  string `json:"data"`
	Dummy []int  `json:"dummy"`
	Term  int    `json:"term"`
}

type Npid struct {
	Handler  NpidHandler `json:"handler"`
	Opt      []int       `json:"opt"`
	Reserved []int       `json:"reserved"`
}

type FobPlayerInfo struct {
	Npid       Npid   `json:"npid"`
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	Ugc        int    `json:"ugc"`
	Xuid       uint64 `json:"xuid"`
}

type Emblem struct {
	Parts []EmblemPart `json:"parts"`
}

type EmblemPart struct {
	BaseColor  int `json:"base_color"`
	FrameColor int `json:"frame_color"`
	PositionX  int `json:"position_x"`
	PositionY  int `json:"position_y"`
	Rotate     int `json:"rotate"`
	Scale      int `json:"scale"`
	TextureTag int `json:"texture_tag"`
}

type FobRank struct {
	Grade int `json:"grade"`
	Rank  int `json:"rank"`
	Score int `json:"score"`
}

type FobPlayerDetailRecord struct {
	Emblem              Emblem    `json:"emblem"`
	Enemy               int       `json:"enemy"`
	Espionage           Espionage `json:"espionage"`
	Follow              int       `json:"follow"`
	Follower            int       `json:"follower"`
	Help                int       `json:"help"`
	Hero                int       `json:"hero"`
	Insurance           int       `json:"insurance"`
	IsSecurityChallenge int       `json:"is_security_challenge"`
	LeagueRank          FobRank   `json:"league_rank"`
	NamePlateID         int       `json:"name_plate_id"`
	Nuclear             int       `json:"nuclear"`
	Online              int       `json:"online"`
	SneakRank           FobRank   `json:"sneak_rank"`
	StaffCount          int       `json:"staff_count"`
}

type OwnerFobRecord struct {
	AttackCount          int                    `json:"attack_count"`
	AttackGmp            int                    `json:"attack_gmp"`
	CaptureNuclear       int                    `json:"capture_nuclear"`
	CaptureResource      CmdMiningResourceEntry `json:"capture_resource"`
	CaptureResourceCount int                    `json:"capture_resource_count"`
	CaptureStaff         int                    `json:"capture_staff"`
	CaptureStaffCount    []int                  `json:"capture_staff_count"`
	DateTime             int                    `json:"date_time"`
	InjuryStaffCount     []int                  `json:"injury_staff_count"`
	LeftHour             int                    `json:"left_hour"`
	NamePlateID          int                    `json:"name_plate_id"`
	Nuclear              int                    `json:"nuclear"`
	ProcessingResource   CmdMiningResourceEntry `json:"processing_resource"`
	StaffCount           []int                  `json:"staff_count"`
	SupportCount         int                    `json:"support_count"`
	SupportedCount       int                    `json:"supported_count"`
	UsableResource       CmdMiningResourceEntry `json:"usable_resource"`
}

type TargetEntry struct {
	AttackerEmblem         Emblem                `json:"attacker_emblem"`
	AttackerEspionage      Espionage             `json:"attacker_espionage"`
	AttackerInfo           FobPlayerInfo         `json:"attacker_info"`
	AttackerSneakRankGrade int                   `json:"attacker_sneak_rank_grade"`
	Cluster                int                   `json:"cluster"`
	IsSneakRestriction     int                   `json:"is_sneak_restriction"`
	IsWin                  int                   `json:"is_win"`
	MotherBaseParam        []MotherBaseParam     `json:"mother_base_param"`
	OwnerDetailRecord      FobPlayerDetailRecord `json:"owner_detail_record"`
	OwnerFobRecord         OwnerFobRecord        `json:"owner_fob_record"`
	OwnerInfo              FobPlayerInfo         `json:"owner_info"`
	SneakMode              int                   `json:"sneak_mode"`
}

type CmdGetFobTargetListResponse struct {
	CryptoType              string               `json:"crypto_type"`
	EnableSecurityChallenge int                  `json:"enable_security_challenge"`
	EspPoint                int                  `json:"esp_point"`
	EventPoint              int                  `json:"event_point"`
	Flowid                  any                  `json:"flowid"`
	FobDeployDamageParam    FobDeployDamageParam `json:"fob_deploy_damage_param"`
	Lose                    int                  `json:"lose"`
	Msgid                   string               `json:"msgid"`
	Result                  string               `json:"result"`
	Rqid                    int                  `json:"rqid"`
	ShieldDate              int                  `json:"shield_date"`
	TargetList              []TargetEntry        `json:"target_list"`
	TargetNum               int                  `json:"target_num"`
	Type                    string               `json:"type"`
	Win                     int                  `json:"win"`
	Xuid                    any                  `json:"xuid"`
}
