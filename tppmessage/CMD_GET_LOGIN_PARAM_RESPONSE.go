package tppmessage

import "github.com/unknown321/fuse/clusterbuildcost"

type CmdGetLoginParamResponse struct {
	ClusterBuildCostsPerCluster []ClusterBuildCostsPerCluster `json:"cluster_build_costs_per_cluster"`
	CryptoType                  string                        `json:"crypto_type"`
	EventPoint                  int                           `json:"event_point"`
	Flowid                      any                           `json:"flowid"`
	FobEventTaskList            FobEventTaskList              `json:"fob_event_task_list"`
	FobEventTaskResultParam     FobEventTaskResultParam       `json:"fob_event_task_result_param"`
	HeroThresholdPoint          int                           `json:"hero_threshold_point"`
	IsAbleToBuyFob4             int                           `json:"is_able_to_buy_fob4"`
	IsStuckRescue               int                           `json:"is_stuck_rescue"`
	Msgid                       string                        `json:"msgid"`
	NotHeroThresholdPoint       int                           `json:"not_hero_threshold_point"`
	OnlineChallengeTask         OnlineChallengeTask           `json:"online_challenge_task"`
	RankingEspiEventIDs         []int                         `json:"ranking_espi_event_ids"`
	RankingPfEventIDs           []int                         `json:"ranking_pf_event_ids"`
	Result                      string                        `json:"result"`
	Rqid                        int                           `json:"rqid"`
	ServerProductParams         []ServerProductParam          `json:"server_product_params"`
	ServerTexts                 []ServerText                  `json:"server_texts"`
	StaffRankBonusRates         []StaffRankBonusRate          `json:"staff_rank_bonus_rates"`
	Xuid                        any                           `json:"xuid"`
}

type ClusterBuildCostsPerGrade struct {
	ClusterBuildCosts []clusterbuildcost.ClusterBuildCost `json:"cluster_build_costs"`
}

type ClusterBuildCostsPerCluster struct {
	ClusterBuildCostsPerGrade []ClusterBuildCostsPerGrade `json:"cluster_build_costs_per_grade"`
}

type FobEventTaskListRewardInfo struct {
	Reward     int `json:"reward"`
	TaskTypeID int `json:"task_type_id"`
	Threshold  int `json:"threshold"`
}

type OneEventTask struct {
	EventID    int                          `json:"event_id"`
	EventSneak []FobEventTaskListRewardInfo `json:"event_sneak"`
}

type FobEventTaskList struct {
	NormalDefense []FobEventTaskListRewardInfo `json:"normal_defense"`
	NormalSneak   []FobEventTaskListRewardInfo `json:"normal_sneak"`
	OneEventTask  []OneEventTask               `json:"one_event_task"`
}

type OneEventParam struct {
	EventID                 int   `json:"event_id"`
	EventSneakClearPointMax int   `json:"event_sneak_clear_point_max"`
	EventSneakClearPointMin int   `json:"event_sneak_clear_point_min"`
	EventSneakSameTimeBonus []int `json:"event_sneak_same_time_bonus"`
}

type FobEventTaskResultParam struct {
	NormalDefenseSameTimeBonus []int           `json:"normal_defense_same_time_bonus"`
	NormalSneakSameTimeBonus   []int           `json:"normal_sneak_same_time_bonus"`
	OneEventParam              []OneEventParam `json:"one_event_param"`
}

type OnlineChallengeTaskReward struct {
	BottomType int `json:"bottom_type"`
	Rate       int `json:"rate"`
	Section    int `json:"section"`
	Type       int `json:"type"`
	Value      int `json:"value"`
}

type OnlineChallengeTaskEntry struct {
	MissionID  int                       `json:"mission_id"`
	Reward     OnlineChallengeTaskReward `json:"reward"`
	Status     int                       `json:"status"`
	TaskTypeID int                       `json:"task_type_id"`
	Threshold  int                       `json:"threshold"`
}

type OnlineChallengeTask struct {
	EndDate  int                        `json:"end_date"`
	TaskList []OnlineChallengeTaskEntry `json:"task_list"`
	Version  int                        `json:"version"`
}

type ServerProductParam struct {
	DevCoin            int `json:"dev_coin"`
	DevGmp             int `json:"dev_gmp"`
	DevItem1           int `json:"dev_item_1"`
	DevItem2           int `json:"dev_item_2"`
	DevPlatlv01        int `json:"dev_platlv01"`
	DevPlatlv02        int `json:"dev_platlv02"`
	DevPlatlv03        int `json:"dev_platlv03"`
	DevPlatlv04        int `json:"dev_platlv04"`
	DevPlatlv05        int `json:"dev_platlv05"`
	DevPlatlv06        int `json:"dev_platlv06"`
	DevPlatlv07        int `json:"dev_platlv07"`
	DevRescount01Value int `json:"dev_rescount01_value"`
	DevRescount02Value int `json:"dev_rescount02_value"`
	DevResource01ID    int `json:"dev_resource01_id"`
	DevResource02ID    int `json:"dev_resource02_id"`
	DevSkil            int `json:"dev_skil"`
	DevSpecial         int `json:"dev_special"`
	DevTime            int `json:"dev_time"`
	ID                 int `json:"id"`
	Open               int `json:"open"`
	Type               int `json:"type"`
	UseGmp             int `json:"use_gmp"`
	UseRescount01Value int `json:"use_rescount01_value"`
	UseRescount02Value int `json:"use_rescount02_value"`
	UseResource01ID    int `json:"use_resource01_id"`
	UseResource02ID    int `json:"use_resource02_id"`
}

type StaffRankBonusRate struct {
	Rates []int `json:"rates"`
}

type ServerText struct {
	Identifier string `json:"identifier"`
	Language   string `json:"language"`
	Text       string `json:"text"`
}
