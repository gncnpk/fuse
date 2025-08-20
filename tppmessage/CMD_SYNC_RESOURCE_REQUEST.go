package tppmessage

type CmdSyncResourceRequest struct {
	CompensateResource []int  `json:"compensate_resource"`
	CumulativeGrade    int    `json:"cumulative_grade"`
	DiffResource1      []int  `json:"diff_resource1"`
	DiffResource2      []int  `json:"diff_resource2"`
	Gmp                int    `json:"gmp"`
	Hero               int    `json:"hero"`
	IsForceBalance     int    `json:"is_force_balance"`
	IsHero             int    `json:"is_hero"`
	IsWallet           int    `json:"is_wallet"`
	Msgid              string `json:"msgid"`
	Rqid               int    `json:"rqid"`
	Version            int    `json:"version"`
}
