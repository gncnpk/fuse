package tppmessage

type CmdMiningResourceEntry struct {
	BioticResource int `json:"biotic_resource"`
	CommonMetal    int `json:"common_metal"`
	FuelResource   int `json:"fuel_resource"`
	MinorMetal     int `json:"minor_metal"`
	PreciousMetal  int `json:"precious_metal"`
}

// FixLimit quick and dirty limits
func (t *CmdMiningResourceEntry) FixLimit(min int, max int) {
	if t.BioticResource < min {
		t.BioticResource = min
	}
	if t.BioticResource > max {
		t.BioticResource = max
	}

	if t.CommonMetal < min {
		t.CommonMetal = min
	}
	if t.CommonMetal > max {
		t.CommonMetal = max
	}

	if t.MinorMetal < min {
		t.MinorMetal = min
	}
	if t.MinorMetal > max {
		t.MinorMetal = max
	}

	if t.PreciousMetal < min {
		t.PreciousMetal = min
	}
	if t.PreciousMetal > max {
		t.PreciousMetal = max
	}

	if t.FuelResource < min {
		t.FuelResource = min
	}
	if t.FuelResource > max {
		t.FuelResource = max
	}
}

type CmdMiningResourceResponse struct {
	AfterMiningResource             CmdMiningResourceEntry `json:"after_mining_resource"`
	AfterMiningResourceProcessing   CmdMiningResourceEntry `json:"after_mining_resource_processing"`
	AfterMiningResourceUsable       CmdMiningResourceEntry `json:"after_mining_resource_usable"`
	AfterProcessResource            CmdMiningResourceEntry `json:"after_process_resource"`
	AfterProcessResourceProcessing  CmdMiningResourceEntry `json:"after_process_resource_processing"`
	AfterProcessResourceUsable      CmdMiningResourceEntry `json:"after_process_resource_usable"`
	BeforeMiningResource            CmdMiningResourceEntry `json:"before_mining_resource"`
	BeforeProcessResource           CmdMiningResourceEntry `json:"before_process_resource"`
	BeforeProcessResourceProcessing CmdMiningResourceEntry `json:"before_process_resource_processing"`
	BeforeProcessResourceUsable     CmdMiningResourceEntry `json:"before_process_resource_usable"`
	BreakTime                       int                    `json:"break_time"`
	CryptoType                      string                 `json:"crypto_type"`
	Flowid                          any                    `json:"flowid"`
	LocalProcessResource            CmdMiningResourceEntry `json:"local_process_resource"`
	Msgid                           string                 `json:"msgid"`
	Result                          string                 `json:"result"`
	Rqid                            int                    `json:"rqid"`
	Xuid                            any                    `json:"xuid"`
}
