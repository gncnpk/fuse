package tppmessage

type CmdGetLoginParamRequest struct {
	Msgid                  string `json:"msgid"`
	Rqid                   int    `json:"rqid"`
	ServerItemPlatformInfo struct {
		PlatformBaseRank int `json:"platform_base_rank"`
	} `json:"server_item_platform_info"`
}
