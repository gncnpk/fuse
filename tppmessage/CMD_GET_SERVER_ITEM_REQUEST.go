package tppmessage

type CmdGetServerItemRequest struct {
	ItemId                 int                    `json:"item_id"`
	Msgid                  string                 `json:"msgid"`
	Rqid                   int                    `json:"rqid"`
	ServerItemPlatformInfo ServerItemPlatformInfo `json:"server_item_platform_info"`
}

type ServerItemPlatformInfo struct {
	PlatformBaseRank       int   `json:"platform_base_rank"`
	SpecialSoldierTypeList []int `json:"special_soldier_type_list"`
}
