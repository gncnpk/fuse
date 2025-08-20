package tppmessage

type CmdGetChallengeTaskRewardsReward struct {
	BottomType int `json:"bottom_type"`
	Rate       int `json:"rate"`
	Section    int `json:"section"`
	Type       int `json:"type"`
	Value      int `json:"value"`
}

type CmdGetChallengeTaskRewardsTask struct {
	Reward CmdGetChallengeTaskRewardsReward `json:"reward"`
	TaskID int                              `json:"task_id"`
}

type CmdGetChallengeTaskRewardsResponse struct {
	CryptoType string                           `json:"crypto_type"`
	Flowid     any                              `json:"flowid"`
	Msgid      string                           `json:"msgid"`
	Result     string                           `json:"result"`
	Rqid       int                              `json:"rqid"`
	TaskCount  int                              `json:"task_count"`
	TaskList   []CmdGetChallengeTaskRewardsTask `json:"task_list"`
	Xuid       any                              `json:"xuid"`
}
