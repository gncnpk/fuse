package tppmessage

type CmdDevelopServerItemRequest struct {
	ID                     int    `json:"id"`
	Msgid                  string `json:"msgid"`
	Rqid                   int    `json:"rqid"`
	ServerItemPlatformInfo struct {
		PlatformBaseRank       int   `json:"platform_base_rank"`
		SpecialSoldierTypeList []int `json:"special_soldier_type_list"`
	} `json:"server_item_platform_info"`
}
