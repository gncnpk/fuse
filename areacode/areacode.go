package areacode

//go:generate go tool stringer -type=EAreaCode
type EAreaCode uint8

// ./chunk0/Assets/tpp/pack/ui/lang/lang_default_data_eng_fpk/Assets/tpp/lang/ui/tpp_common.eng.lng2.xml
const (
	CentralIndianRidge       EAreaCode = 0
	MidAtlanticRidge         EAreaCode = 10
	EastOfTheHawaiianIslands EAreaCode = 20
	SouthAtlanticOcean       EAreaCode = 30
	IndianOcean              EAreaCode = 40
	NorthPacificOcean        EAreaCode = 70
	SouthPacificOcean        EAreaCode = 80
	NorthAtlanticOcean       EAreaCode = 90
)

var validAreaCodes = map[EAreaCode]bool{
	CentralIndianRidge:       true,
	MidAtlanticRidge:         true,
	EastOfTheHawaiianIslands: true,
	SouthAtlanticOcean:       true,
	IndianOcean:              true,
	NorthPacificOcean:        true,
	SouthPacificOcean:        true,
	NorthAtlanticOcean:       true,
}

func IsValid(code int) bool {
	_, exists := validAreaCodes[EAreaCode(code)]
	return exists
}
