package tppmessage

type CmdSendMissionResultRequest struct {
	MissionID  int    `json:"mission_id"`
	Msgid      string `json:"msgid"`
	Rqid       int    `json:"rqid"`
	SendRecord struct {
		BattleGearMove int   `json:"battle_gear_move"`
		Cbox           []int `json:"cbox"`
		CrackClimb     int   `json:"crack_climb"`
		CrawlMove      int   `json:"crawl_move"`
		EspionageRadio int   `json:"espionage_radio"`
		GarbageBox     int   `json:"garbage_box"`
		HorseMove      int   `json:"horse_move"`
		OptionalRadio  int   `json:"optional_Radio"`
		SquatMove      int   `json:"squat_move"`
		StandMove      int   `json:"stand_move"`
		Suit           []int `json:"suit"`
		Toilet         int   `json:"toilet"`
		VehicleMove    int   `json:"vehicle_move"`
		WalkerGearMove int   `json:"walker_gear_move"`
	} `json:"send_record"`
}
