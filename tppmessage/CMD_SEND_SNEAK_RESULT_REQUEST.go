package tppmessage

type CmdSendSneakResultRequest struct {
	CaptureNuclear   int `json:"capture_nuclear"`
	CapturePlacement struct {
		EmplacementGunEast int `json:"emplacement_gun_east"`
		EmplacementGunWest int `json:"emplacement_gun_west"`
		GatlingGun         int `json:"gatling_gun"`
		GatlingGunEast     int `json:"gatling_gun_east"`
		GatlingGunWest     int `json:"gatling_gun_west"`
		MortarNormal       int `json:"mortar_normal"`
	} `json:"capture_placement"`
	CapturePlayerSoldierNum int `json:"capture_player_soldier_num"`
	CaptureResource         struct {
		BioticResource int `json:"biotic_resource"`
		CommonMetal    int `json:"common_metal"`
		FuelResource   int `json:"fuel_resource"`
		MinorMetal     int `json:"minor_metal"`
		PreciousMetal  int `json:"precious_metal"`
	} `json:"capture_resource"`
	CaptureSoldierCount          []int `json:"capture_soldier_count"`
	CaptureSoldierNum            int   `json:"capture_soldier_num"`
	CaptureSupportSoldierNum     int   `json:"capture_support_soldier_num"`
	ClearedPlantCount            int   `json:"cleared_plant_count"`
	CountOfNeutralizeIntruder    int   `json:"count_of_neutralize_intruder"`
	CountOfNeutralizedByIntruder int   `json:"count_of_neutralized_by_intruder"`
	DamagePoint                  int   `json:"damage_point"`
	DestroyPlacement             struct {
		EmplacementGunEast int `json:"emplacement_gun_east"`
		EmplacementGunWest int `json:"emplacement_gun_west"`
		GatlingGun         int `json:"gatling_gun"`
		GatlingGunEast     int `json:"gatling_gun_east"`
		GatlingGunWest     int `json:"gatling_gun_west"`
		MortarNormal       int `json:"mortar_normal"`
	} `json:"destroy_placement"`
	Event struct {
		AttackerInfo struct {
			Npid struct {
				Handler struct {
					Data string `json:"data"`
					Term int    `json:"term"`
				} `json:"handler"`
			} `json:"npid"`
			PlayerID   int    `json:"player_id"`
			PlayerName string `json:"player_name"`
			Ugc        int    `json:"ugc"`
			Xuid       int    `json:"xuid"`
		} `json:"attacker_info"`
		AttackerLeagueGrade int `json:"attacker_league_grade"`
		AttackerSneakGrade  int `json:"attacker_sneak_grade"`
		CaptureNuclear      int `json:"capture_nuclear"`
		CaptureResource     struct {
			BioticResource int `json:"biotic_resource"`
			CommonMetal    int `json:"common_metal"`
			FuelResource   int `json:"fuel_resource"`
			MinorMetal     int `json:"minor_metal"`
			PreciousMetal  int `json:"precious_metal"`
		} `json:"capture_resource"`
		Cluster    int    `json:"cluster"`
		Data       string `json:"data"`
		Gmp        int    `json:"gmp"`
		IsWin      int    `json:"is_win"`
		LayoutCode int    `json:"layout_code"`
		PositionX  int    `json:"position_x"`
		PositionZ  int    `json:"position_z"`
		RegistDate int    `json:"regist_date"`
		RotateY    int    `json:"rotate_y"`
		Size       int    `json:"size"`
	} `json:"event"`
	EventPoint              int    `json:"event_point"`
	EventVersion            int    `json:"event_version"`
	HighRank                int    `json:"high_rank"`
	InjureSoldierNum        int    `json:"injure_soldier_num"`
	InjureSupportSoldierNum int    `json:"injure_support_soldier_num"`
	IsEvent                 int    `json:"is_event"`
	IsGoal                  int    `json:"is_goal"`
	IsPerfectStealth        int    `json:"is_perfect_stealth"`
	IsPlus                  int    `json:"is_plus"`
	IsSneak                 int    `json:"is_sneak"`
	IsSupporter             int    `json:"is_supporter"`
	KillSoldierNum          int    `json:"kill_soldier_num"`
	KillSupportSoldierNum   int    `json:"kill_support_soldier_num"`
	MissionTaskCompleteBits int    `json:"mission_task_complete_bits"`
	Mode                    string `json:"mode"`
	MotherBaseID            int    `json:"mother_base_id"`
	Msgid                   string `json:"msgid"`
	OpenWormhole            int    `json:"open_wormhole"`
	RecoverResource         struct {
		BioticResource int `json:"biotic_resource"`
		CommonMetal    int `json:"common_metal"`
		FuelResource   int `json:"fuel_resource"`
		MinorMetal     int `json:"minor_metal"`
		PreciousMetal  int `json:"precious_metal"`
	} `json:"recover_resource"`
	ResultType        int    `json:"result_type"`
	RetaliatePoint    int    `json:"retaliate_point"`
	RetaliateWormhole int    `json:"retaliate_wormhole"`
	Rqid              int    `json:"rqid"`
	SneakPoint        int    `json:"sneak_point"`
	SneakResult       string `json:"sneak_result"`
	Version           int    `json:"version"`
}
