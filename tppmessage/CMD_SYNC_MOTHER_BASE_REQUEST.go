package tppmessage

type CmdSyncMotherBaseRequest struct {
	EquipFlag       []int  `json:"equip_flag"`
	EquipGrade      []int  `json:"equip_grade"`
	Flag            string `json:"flag"`
	InvalidFob      int    `json:"invalid_fob"`
	LocalBaseParam  []int  `json:"local_base_param"`
	LocalBaseTime   []int  `json:"local_base_time"`
	MotherBaseNum   int    `json:"mother_base_num"`
	MotherBaseParam []struct {
		AreaID       int `json:"area_id"`
		ClusterParam []struct {
			Build           int `json:"build"`
			ClusterSecurity int `json:"cluster_security"`
			Common1Security struct {
				Antitheft                  int                    `json:"antitheft"`
				Camera                     int                    `json:"camera"`
				CautionArea                int                    `json:"caution_area"`
				Decoy                      int                    `json:"decoy"`
				IrSensor                   int                    `json:"ir_sensor"`
				Mine                       int                    `json:"mine"`
				Soldier                    int                    `json:"soldier"`
				Uav                        int                    `json:"uav"`
				VoluntaryCoordCameraCount  int                    `json:"voluntary_coord_camera_count"`
				VoluntaryCoordCameraParams []VoluntaryCoordParams `json:"voluntary_coord_camera_params"`
				VoluntaryCoordMineCount    int                    `json:"voluntary_coord_mine_count"`
				VoluntaryCoordMineParams   []VoluntaryCoordParams `json:"voluntary_coord_mine_params"`
			} `json:"common1_security"`
			Common2Security struct {
				Antitheft                  int                    `json:"antitheft"`
				Camera                     int                    `json:"camera"`
				CautionArea                int                    `json:"caution_area"`
				Decoy                      int                    `json:"decoy"`
				IrSensor                   int                    `json:"ir_sensor"`
				Mine                       int                    `json:"mine"`
				Soldier                    int                    `json:"soldier"`
				Uav                        int                    `json:"uav"`
				VoluntaryCoordCameraCount  int                    `json:"voluntary_coord_camera_count"`
				VoluntaryCoordCameraParams []VoluntaryCoordParams `json:"voluntary_coord_camera_params"`
				VoluntaryCoordMineCount    int                    `json:"voluntary_coord_mine_count"`
				VoluntaryCoordMineParams   []VoluntaryCoordParams `json:"voluntary_coord_mine_params"`
			} `json:"common2_security"`
			Common3Security struct {
				Antitheft                  int                    `json:"antitheft"`
				Camera                     int                    `json:"camera"`
				CautionArea                int                    `json:"caution_area"`
				Decoy                      int                    `json:"decoy"`
				IrSensor                   int                    `json:"ir_sensor"`
				Mine                       int                    `json:"mine"`
				Soldier                    int                    `json:"soldier"`
				Uav                        int                    `json:"uav"`
				VoluntaryCoordCameraCount  int                    `json:"voluntary_coord_camera_count"`
				VoluntaryCoordCameraParams []VoluntaryCoordParams `json:"voluntary_coord_camera_params"`
				VoluntaryCoordMineCount    int                    `json:"voluntary_coord_mine_count"`
				VoluntaryCoordMineParams   []VoluntaryCoordParams `json:"voluntary_coord_mine_params"`
			} `json:"common3_security"`
			SoldierRank    int `json:"soldier_rank"`
			UniqueSecurity struct {
				Antitheft                  int                    `json:"antitheft"`
				Camera                     int                    `json:"camera"`
				CautionArea                int                    `json:"caution_area"`
				Decoy                      int                    `json:"decoy"`
				IrSensor                   int                    `json:"ir_sensor"`
				Mine                       int                    `json:"mine"`
				Soldier                    int                    `json:"soldier"`
				Uav                        int                    `json:"uav"`
				VoluntaryCoordCameraCount  int                    `json:"voluntary_coord_camera_count"`
				VoluntaryCoordCameraParams []VoluntaryCoordParams `json:"voluntary_coord_camera_params"`
				VoluntaryCoordMineCount    int                    `json:"voluntary_coord_mine_count"`
				VoluntaryCoordMineParams   []VoluntaryCoordParams `json:"voluntary_coord_mine_params"`
			} `json:"unique_security"`
		} `json:"cluster_param"`
		ConstructParam int `json:"construct_param"`
		FobIndex       int `json:"fob_index"`
		MotherBaseID   int `json:"mother_base_id"`
		PlatformCount  int `json:"platform_count"`
		Price          int `json:"price"`
		SecurityRank   int `json:"security_rank"`
	} `json:"mother_base_param"`
	Msgid        string `json:"msgid"`
	NamePlateID  int    `json:"name_plate_id"`
	PfSkillStaff struct {
		AllStaffNum             int `json:"all_staff_num"`
		Defender1Num            int `json:"defender1_num"`
		Defender2Num            int `json:"defender2_num"`
		Defender3Num            int `json:"defender3_num"`
		InterceptorMissile1Num  int `json:"interceptor_missile1_num"`
		InterceptorMissile2Num  int `json:"interceptor_missile2_num"`
		InterceptorMissile3Num  int `json:"interceptor_missile3_num"`
		LiquidCarbonMissile1Num int `json:"liquid_carbon_missile1_num"`
		LiquidCarbonMissile2Num int `json:"liquid_carbon_missile2_num"`
		LiquidCarbonMissile3Num int `json:"liquid_carbon_missile3_num"`
		Medic1Num               int `json:"medic1_num"`
		Medic2Num               int `json:"medic2_num"`
		Medic3Num               int `json:"medic3_num"`
		Ranger1Num              int `json:"ranger1_num"`
		Ranger2Num              int `json:"ranger2_num"`
		Ranger3Num              int `json:"ranger3_num"`
		Sentry1Num              int `json:"sentry1_num"`
		Sentry2Num              int `json:"sentry2_num"`
		Sentry3Num              int `json:"sentry3_num"`
	} `json:"pf_skill_staff"`
	PickupOpen    int   `json:"pickup_open"`
	Rqid          int   `json:"rqid"`
	SectionOpen   int   `json:"section_open"`
	SecurityLevel []int `json:"security_level"`
	TapeFlag      []int `json:"tape_flag"`
	Version       int   `json:"version"`
}
