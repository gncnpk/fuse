package tppmessage

type CmdCheckServerItemCorrectRequest struct {
	ItemList    []int  `json:"item_list"`
	ItemListNum int    `json:"item_list_num"`
	Msgid       string `json:"msgid"`
	Rqid        int    `json:"rqid"`
}
