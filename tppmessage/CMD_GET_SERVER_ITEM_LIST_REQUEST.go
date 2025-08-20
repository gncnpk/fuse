package tppmessage

type CmdGetServerItemListRequest struct {
	ListMaxCount           int    `json:"list_max_count"`
	Msgid                  string `json:"msgid"`
	Rqid                   int    `json:"rqid"`
	ServerItemPlatformInfo struct {
		PlatformBaseRank       int   `json:"platform_base_rank"`
		SpecialSoldierTypeList []int `json:"special_soldier_type_list"`
	} `json:"server_item_platform_info"`
}
