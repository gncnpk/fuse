package tppmessage

type FobSoldier struct {
	Header       uint64 `json:"header"`
	Seed         int    `json:"seed"`
	StatusNoSync int    `json:"status_no_sync"`
	StatusSync   int    `json:"status_sync"`
}

type FobSneakMotherBaseStageParam struct {
	Build              []int                  `json:"build"`
	ClusterParam       ClusterParam           `json:"cluster_param"`
	ConstructParam     int                    `json:"construct_param"`
	EquipGrade         []int                  `json:"equip_grade"`
	FobIndex           int                    `json:"fob_index"`
	MotherBaseID       int                    `json:"mother_base_id"`
	Nuclear            int                    `json:"nuclear"`
	OwnerPlayerID      int                    `json:"owner_player_id"`
	Placement          FobWeaponPlacement     `json:"placement"`
	Platform           int                    `json:"platform"`
	ProcessingResource CmdMiningResourceEntry `json:"processing_resource"`
	SectionLevel       FobSectionStaff        `json:"section_level"`
	SecurityLevel      []int                  `json:"security_level"`
	UsableResource     CmdMiningResourceEntry `json:"usable_resource"`
}

type CmdSneakMotherBaseResponse struct {
	CryptoType           string                       `json:"crypto_type"`
	DamageParam          []any                        `json:"damage_param"`
	EventFobParams       []int                        `json:"event_fob_params"`
	Flowid               any                          `json:"flowid"`
	FobDeployDamageParam FobDeployDamageParam         `json:"fob_deploy_damage_param"`
	IsEvent              int                          `json:"is_event"`
	IsSecurityContract   int                          `json:"is_security_contract"`
	Msgid                string                       `json:"msgid"`
	OwnerGmp             int                          `json:"owner_gmp"`
	RecoverResource      CmdMiningResourceEntry       `json:"recover_resource"`
	RecoverSoldier       []any                        `json:"recover_soldier"`
	RecoverSoldierCount  []int                        `json:"recover_soldier_count"`
	RecoverSoldierNum    int                          `json:"recover_soldier_num"`
	Result               string                       `json:"result"`
	RewardID             int                          `json:"reward_id"`
	RewardSoldier        []FobSoldier                 `json:"reward_soldier"`
	RewardSoldierNum     int                          `json:"reward_soldier_num"`
	RewardSoldierRank    int                          `json:"reward_soldier_rank"`
	RewardSoldierType    int                          `json:"reward_soldier_type"`
	Rqid                 int                          `json:"rqid"`
	SecuritySoldier      []FobSoldier                 `json:"security_soldier"`
	SecuritySoldierNum   int                          `json:"security_soldier_num"`
	SecuritySoldierRank  int                          `json:"security_soldier_rank"`
	StageParam           FobSneakMotherBaseStageParam `json:"stage_param"`
	WormholePlayerID     int                          `json:"wormhole_player_id"`
	Xuid                 any                          `json:"xuid"`
}
