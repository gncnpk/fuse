package tppmessage

type CmdSyncEmblemRequest struct {
	Emblem Emblem `json:"emblem"`
	Msgid  string `json:"msgid"`
	Rqid   int    `json:"rqid"`
}
