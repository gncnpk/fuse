package tppmessage

type CmdSyncSoldierBinRequest struct {
	Flag      string `json:"flag"`
	ForceSync int    `json:"force_sync"`
	Msgid     string `json:"msgid"`
	Rqid      int    `json:"rqid"`
	Section   struct {
		Base     int `json:"base"`
		Combat   int `json:"combat"`
		Develop  int `json:"develop"`
		Medical  int `json:"medical"`
		Security int `json:"security"`
		Spy      int `json:"spy"`
		Suport   int `json:"suport"`
	} `json:"section"`
	SectionSoldier struct {
		Base     int `json:"base"`
		Combat   int `json:"combat"`
		Develop  int `json:"develop"`
		Medical  int `json:"medical"`
		Security int `json:"security"`
		Spy      int `json:"spy"`
		Suport   int `json:"suport"`
	} `json:"section_soldier"`
	SoldierNum   int    `json:"soldier_num"`
	SoldierParam string `json:"soldier_param"`
	Version      int    `json:"version"`
}
