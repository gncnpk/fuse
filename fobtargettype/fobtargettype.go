package fobtargettype

//go:generate go tool stringer -type=EFOBTargetType
type EFOBTargetType uint8

const (
	Invalid EFOBTargetType = iota
	PICKUP                           // infiltration targets, pfs of equal grade
	PICKUP_HIGH                      // infiltration targets, high-ranking pfs
	ENEMY                            // retaliation targets
	FR_ENEMY                         // indirect retaliation targets
	NUCLEAR                          // nuclear-equipped targets
	TRIAL                            // training/visit destination
	EVENTS                           // events
	INJURY                           // intruder
	CHALLENGE                        // security challenge
	DEPLOYED                         // fob unit deployed list
	EMERGENCY                        // intruder alert
)
