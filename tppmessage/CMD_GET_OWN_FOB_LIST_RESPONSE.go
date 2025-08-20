package tppmessage

type VoluntaryCoordParams struct {
	PlacedIndex int `json:"placed_index"`
	PositionX   int `json:"position_x"`
	PositionY   int `json:"position_y"`
	PositionZ   int `json:"position_z"`
	RotationW   int `json:"rotation_w"`
	RotationX   int `json:"rotation_x"`
	RotationY   int `json:"rotation_y"`
	RotationZ   int `json:"rotation_z"`
}

type ClusterParamSecurity struct {
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
}

type ClusterParam struct {
	Build           int                  `json:"build"`
	ClusterSecurity int                  `json:"cluster_security"`
	Common1Security ClusterParamSecurity `json:"common1_security"`
	Common2Security ClusterParamSecurity `json:"common2_security"`
	Common3Security ClusterParamSecurity `json:"common3_security"`
	SoldierRank     int                  `json:"soldier_rank"`
	UniqueSecurity  ClusterParamSecurity `json:"unique_security"`
}

type MotherBaseParam struct {
	AreaID         int            `json:"area_id"`
	ClusterParam   []ClusterParam `json:"cluster_param"`
	ConstructParam int            `json:"construct_param"`
	FobIndex       int            `json:"fob_index"`
	MotherBaseID   int            `json:"mother_base_id"`
	PlatformCount  int            `json:"platform_count"`
	Price          int            `json:"price"`
	SecurityRank   int            `json:"security_rank"`
}

type CmdGetOwnFobListResponse struct {
	CryptoType              string            `json:"crypto_type"`
	EnableSecurityChallenge int               `json:"enable_security_challenge"`
	Flowid                  any               `json:"flowid"`
	Fob                     []MotherBaseParam `json:"fob"`
	Msgid                   string            `json:"msgid"`
	Result                  string            `json:"result"`
	Rqid                    int               `json:"rqid"`
	Xuid                    any               `json:"xuid"`
}
