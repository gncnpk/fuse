package tppmessage

type CmdSyncLoadoutRequest struct {
	Loadout struct {
		Hand struct {
			ID        int   `json:"id"`
			LevelList []int `json:"level_list"`
		} `json:"hand"`
		ItemLevelList []int `json:"item_level_list"`
		ItemList      []int `json:"item_list"`
		Primary1      struct {
			ChimeraDesc struct {
				Color     int   `json:"color"`
				Paint     int   `json:"paint"`
				PartsList []int `json:"parts_list"`
			} `json:"chimera_desc"`
			ChimeraSlot int `json:"chimera_slot"`
			ID          int `json:"id"`
			IsChimera   int `json:"is_chimera"`
		} `json:"primary1"`
		Primary2 struct {
			ChimeraDesc struct {
				Color     int   `json:"color"`
				Paint     int   `json:"paint"`
				PartsList []int `json:"parts_list"`
			} `json:"chimera_desc"`
			ChimeraSlot int `json:"chimera_slot"`
			ID          int `json:"id"`
			IsChimera   int `json:"is_chimera"`
		} `json:"primary2"`
		Secondary struct {
			ChimeraDesc struct {
				Color     int   `json:"color"`
				Paint     int   `json:"paint"`
				PartsList []int `json:"parts_list"`
			} `json:"chimera_desc"`
			ChimeraSlot int `json:"chimera_slot"`
			ID          int `json:"id"`
			IsChimera   int `json:"is_chimera"`
		} `json:"secondary"`
		Suit struct {
			Camo  int `json:"camo"`
			Face  int `json:"face"`
			Level int `json:"level"`
			Parts int `json:"parts"`
		} `json:"suit"`
		SupportList []int `json:"support_list"`
	} `json:"loadout"`
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}
