package tppmessage

type CmdSyncResourceResponse struct {
	CryptoType    string `json:"crypto_type"`
	DiffGmp       int    `json:"diff_gmp"`
	DiffResource1 []int  `json:"diff_resource1"`
	DiffResource2 []int  `json:"diff_resource2"`

	/* also known as `before_process_resource_usable` from CMD_MINING_RESOURCE
		- fuel
		- biotic
		- common metal
		- minor metal
		- precious metal
	    - ???
	*/
	FixResource1 []int `json:"fix_resource1"`

	/* also known as `before_process_resource_processing` from CMD_MINING_RESOURCE
		- fuel
		- biotic
		- common metal
		- minor metal
		- precious metal
	    - ???
	*/
	FixResource2 []int  `json:"fix_resource2"`
	Flowid       any    `json:"flowid"`
	InjuryGmp    int    `json:"injury_gmp"`
	InsuranceGmp int    `json:"insurance_gmp"`
	LoadoutGmp   int    `json:"loadout_gmp"`
	LocalGmp     int    `json:"local_gmp"`
	Msgid        string `json:"msgid"`
	Result       string `json:"result"`
	Rqid         int    `json:"rqid"`
	ServerGmp    int    `json:"server_gmp"`
	Version      int    `json:"version"`
	Xuid         any    `json:"xuid"`
}
