package tppmessage

type CmdGetLeagueResultResponse struct {
	CryptoType string `json:"crypto_type"`
	Flowid     any    `json:"flowid"`
	Info       struct {
		Begin        int `json:"begin"`
		Current      int `json:"current"`
		DayBattle    int `json:"day_battle"`
		End          int `json:"end"`
		HistoryCount int `json:"history_count"`
		MatchHistory []struct {
			AttackDurability  int `json:"attack_durability"`
			AttackItem        int `json:"attack_item"`
			AttackLevel       int `json:"attack_level"`
			AttackNuclear     int `json:"attack_nuclear"`
			AttackPid         int `json:"attack_pid"`
			AttackPoint       int `json:"attack_point"`
			AttackStaff       int `json:"attack_staff"`
			CumulativeGrade   int `json:"cumulative_grade"`
			DefenceDurability int `json:"defence_durability"`
			DefenceItem       int `json:"defence_item"`
			DefenceLevel      int `json:"defence_level"`
			DefenceNuclear    int `json:"defence_nuclear"`
			DefencePid        int `json:"defence_pid"`
			DefencePoint      int `json:"defence_point"`
			DefenceStaff      int `json:"defence_staff"`
			MatchDate         int `json:"match_date"`
			Result            int `json:"result"`
			Section           int `json:"section"`
			SecurityRank      int `json:"security_rank"`
			Weather           int `json:"weather"`
			WinPoint          int `json:"win_point"`
		} `json:"match_history"`
		Next        int `json:"next"`
		PlayerCount int `json:"player_count"`
		PlayerInfo  []struct {
			AttackDurability  int `json:"attack_durability"`
			AttackItem        int `json:"attack_item"`
			AttackLevel       int `json:"attack_level"`
			AttackLose        int `json:"attack_lose"`
			AttackNuclear     int `json:"attack_nuclear"`
			AttackPoint       int `json:"attack_point"`
			AttackStaff       int `json:"attack_staff"`
			AttackWin         int `json:"attack_win"`
			ConbatPoint       int `json:"conbat_point"`
			CumulativeGrade   int `json:"cumulative_grade"`
			DefenceDurability int `json:"defence_durability"`
			DefenceItem       int `json:"defence_item"`
			DefenceLevel      int `json:"defence_level"`
			DefenceNuclear    int `json:"defence_nuclear"`
			DefencePoint      int `json:"defence_point"`
			DefenceStaff      int `json:"defence_staff"`
			DefenseLose       int `json:"defense_lose"`
			DefenseWin        int `json:"defense_win"`
			Lose              int `json:"lose"`
			MotherBaseParam   []struct {
				AreaID         int   `json:"area_id"`
				ClusterParam   []any `json:"cluster_param"`
				ConstructParam int   `json:"construct_param"`
				FobIndex       int   `json:"fob_index"`
				MotherBaseID   int   `json:"mother_base_id"`
				PlatformCount  int   `json:"platform_count"`
				Price          int   `json:"price"`
				SecurityRank   int   `json:"security_rank"`
			} `json:"mother_base_param"`
			NarrowLose         int `json:"narrow_lose"`
			NarrowWin          int `json:"narrow_win"`
			PlannedAttackItem  int `json:"planned_attack_item"`
			PlannedDefenceItem int `json:"planned_defence_item"`
			PlayerDetailRecord struct {
				Emblem struct {
					Parts []struct {
						BaseColor  int `json:"base_color"`
						FrameColor int `json:"frame_color"`
						PositionX  int `json:"position_x"`
						PositionY  int `json:"position_y"`
						Rotate     int `json:"rotate"`
						Scale      int `json:"scale"`
						TextureTag int `json:"texture_tag"`
					} `json:"parts"`
				} `json:"emblem"`
				Enemy     int `json:"enemy"`
				Espionage struct {
					Lose    int `json:"lose"`
					Score   int `json:"score"`
					Section int `json:"section"`
					Win     int `json:"win"`
				} `json:"espionage"`
				Follow              int `json:"follow"`
				Follower            int `json:"follower"`
				Help                int `json:"help"`
				Hero                int `json:"hero"`
				Insurance           int `json:"insurance"`
				IsSecurityChallenge int `json:"is_security_challenge"`
				LeagueRank          struct {
					Grade int `json:"grade"`
					Rank  int `json:"rank"`
					Score int `json:"score"`
				} `json:"league_rank"`
				NamePlateID int `json:"name_plate_id"`
				Nuclear     int `json:"nuclear"`
				Online      int `json:"online"`
				SneakRank   struct {
					Grade int `json:"grade"`
					Rank  int `json:"rank"`
					Score int `json:"score"`
				} `json:"sneak_rank"`
				StaffCount int `json:"staff_count"`
			} `json:"player_detail_record"`
			PlayerInfo struct {
				Npid struct {
					Handler struct {
						Data  string `json:"data"`
						Dummy []int  `json:"dummy"`
						Term  int    `json:"term"`
					} `json:"handler"`
					Opt      []int `json:"opt"`
					Reserved []int `json:"reserved"`
				} `json:"npid"`
				PlayerID   int    `json:"player_id"`
				PlayerName string `json:"player_name"`
				Ugc        int    `json:"ugc"`
				Xuid       int    `json:"xuid"`
			} `json:"player_info"`
			Rank          int   `json:"rank"`
			ResultHistory []int `json:"result_history"`
			SecurityRank  int   `json:"security_rank"`
			Win           int   `json:"win"`
			WinningPoint  int   `json:"winning_point"`
		} `json:"player_info"`
		Point   int `json:"point"`
		Section int `json:"section"`
	} `json:"info"`
	Msgid  string `json:"msgid"`
	Result string `json:"result"`
	Rqid   int    `json:"rqid"`
	Xuid   any    `json:"xuid"`
}
