package weaponrange

//go:generate go tool stringer -type=EWeaponRange
type EWeaponRange uint8

const (
	Close EWeaponRange = iota
	Mid
	Long
	End
)

func IsValid(b byte) bool {
	if b < byte(Close) || b >= byte(End) {
		return false
	}

	return true
}
