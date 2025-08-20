package tppmessage

type CmdGetUrllistRequest struct {
	Lang   string `json:"lang"`
	Msgid  string `json:"msgid"`
	Region string `json:"region"`
	Rqid   int    `json:"rqid"`
}
