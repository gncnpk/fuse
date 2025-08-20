package tppmessage

type CmdSneakMotherBaseRequest struct {
	FobIndex            int    `json:"fob_index"`
	IsEvent             int    `json:"is_event"`
	IsPlus              int    `json:"is_plus"`
	IsSecurityChallenge int    `json:"is_security_challenge"`
	IsSneak             int    `json:"is_sneak"`
	Mode                string `json:"mode"`
	MotherBaseID        int    `json:"mother_base_id"` // id 2 is used for fob events?
	Msgid               string `json:"msgid"`
	Platform            int    `json:"platform"`
	PlayerID            int    `json:"player_id"` // always 0
	Rqid                int    `json:"rqid"`
	WormholePlayerID    int    `json:"wormhole_player_id"`
	Xnkey               string `json:"xnkey"`
	Xnkid               string `json:"xnkid"`
}
