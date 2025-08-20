package tppmessage

type CmdGetChallengeTaskTargetValuesResponse struct {
	CryptoType                          string `json:"crypto_type"`
	EspionageRatingGrade                int    `json:"espionage_rating_grade"`
	Flowid                              any    `json:"flowid"`
	FobDefenseSuccessCount              int    `json:"fob_defense_success_count"`
	FobDeployToSupportersEmergencyCount int    `json:"fob_deploy_to_supporters_emergency_count"`
	FobSneakCount                       int    `json:"fob_sneak_count"`
	FobSneakSuccessCount                int    `json:"fob_sneak_success_count"`
	FobSupportingUserCount              int    `json:"fob_supporting_user_count"`
	Msgid                               string `json:"msgid"`
	PfRatingDefenseForce                int    `json:"pf_rating_defense_force"`
	PfRatingDefenseLife                 int    `json:"pf_rating_defense_life"`
	PfRatingOffenceForce                int    `json:"pf_rating_offence_force"`
	PfRatingOffenceLife                 int    `json:"pf_rating_offence_life"`
	PfRatingRank                        int    `json:"pf_rating_rank"`
	Result                              string `json:"result"`
	Rqid                                int    `json:"rqid"`
	TotalDevelopmentGrade               int    `json:"total_development_grade"`
	TotalFobSecurityLevel               int    `json:"total_fob_security_level"`
	Xuid                                any    `json:"xuid"`
}
