package tppmessage

type CmdGetInfoListEntry struct {
	Date       int    `json:"date"`
	Important  string `json:"important"`
	InfoID     int    `json:"info_id"`
	MesBody    string `json:"mes_body"`
	MesSubject string `json:"mes_subject"`
}

type CmdGetInformationlist2Response struct {
	CryptoType string                `json:"crypto_type"`
	Flowid     any                   `json:"flowid"`
	InfoList   []CmdGetInfoListEntry `json:"info_list"`
	InfoNum    int                   `json:"info_num"`
	Msgid      string                `json:"msgid"`
	Result     string                `json:"result"`
	Rqid       int                   `json:"rqid"`
	Xuid       any                   `json:"xuid"`
}
