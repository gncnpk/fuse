package tppmessage

type CmdGetCombatDeployListResponse struct {
	CryptoType  string `json:"crypto_type"`
	Flowid      any    `json:"flowid"`
	MissionList []struct {
		ArmoredMax    int `json:"armored_max"`
		ArmoredMin    int `json:"armored_min"`
		BattleGear    int `json:"battle_gear"`
		CarMax        int `json:"car_max"`
		CarMin        int `json:"car_min"`
		Category      int `json:"category"`
		CombatCount   int `json:"combat_count"`
		CombatRank    int `json:"combat_rank"`
		DeadRate      int `json:"dead_rate"`
		IsCampaign    int `json:"is_campaign"`
		Latitude      int `json:"latitude"`
		Longitude     int `json:"longitude"`
		MaxDeadRate   int `json:"max_dead_rate"`
		MaxWinRate    int `json:"max_win_rate"`
		MinDeadRate   int `json:"min_dead_rate"`
		MinWinRate    int `json:"min_win_rate"`
		MissionID     int `json:"mission_id"`
		NameKey       int `json:"name_key"`
		PrimaryReward []struct {
			BottomType int `json:"bottom_type"`
			MechaType  int `json:"mecha_type"`
			Rate       int `json:"rate"`
			Section    int `json:"section"`
			Type       int `json:"type"`
			Value      int `json:"value"`
		} `json:"primary_reward"`
		Reward       int `json:"reward"`
		Section      int `json:"section"`
		SectionCount int `json:"section_count"`
		SectionRank  int `json:"section_rank"`
		Seed         int `json:"seed"`
		ServerTextID int `json:"server_text_id"`
		TankMax      int `json:"tank_max"`
		TankMin      int `json:"tank_min"`
		Team         struct {
			Armored          int `json:"armored"`
			BattleGear       int `json:"battle_gear"`
			Car              int `json:"car"`
			CombatCount      int `json:"combat_count"`
			CombatRankBottom int `json:"combat_rank_bottom"`
			CombatRankTop    int `json:"combat_rank_top"`
			DeadRate         int `json:"dead_rate"`
			IsValid          int `json:"is_valid"`
			MissionID        int `json:"mission_id"`
			Seed             int `json:"seed"`
			StaffPower       int `json:"staff_power"`
			SubCount         int `json:"sub_count"`
			SubRankBottom    int `json:"sub_rank_bottom"`
			SubRankTop       int `json:"sub_rank_top"`
			Tank             int `json:"tank"`
			TeamID           int `json:"team_id"`
			TeamPower        int `json:"team_power"`
			Time             int `json:"time"`
			Truck            int `json:"truck"`
			WalkerGear       int `json:"walker_gear"`
			WinRate          int `json:"win_rate"`
		} `json:"team"`
		Time          int `json:"time"`
		TimeRandom    int `json:"time_random"`
		TruckMax      int `json:"truck_max"`
		TruckMin      int `json:"truck_min"`
		WalkerGearMax int `json:"walker_gear_max"`
		WalkerGearMin int `json:"walker_gear_min"`
		WinRate       int `json:"win_rate"`
	} `json:"mission_list"`
	MissionNum int    `json:"mission_num"`
	Msgid      string `json:"msgid"`
	Result     string `json:"result"`
	Rqid       int    `json:"rqid"`
	Xuid       any    `json:"xuid"`
}
