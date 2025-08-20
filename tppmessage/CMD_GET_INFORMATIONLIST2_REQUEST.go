package tppmessage

type CmdGetInformationlist2Request struct {
	IsMgo  int    `json:"is_mgo"`
	Lang   string `json:"lang"`
	Msgid  string `json:"msgid"`
	Region string `json:"region"`
	Rqid   int    `json:"rqid"`
}
