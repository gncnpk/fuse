package tppmessage

type CmdDeployMissionRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
	Team  struct {
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
}
