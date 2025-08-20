package guardrank

//go:generate go tool stringer -type=EGuardRank
type EGuardRank uint8

const (
	E EGuardRank = iota
	D
	C
	B
	A
	APlus
	APlusPlus
	S
	SPlus
	SPlusPlus
)

func IsValid(i byte) bool {
	if i < byte(E) || i > byte(SPlusPlus) {
		return false
	}

	return true
}
