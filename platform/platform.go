package platform
//go:generate go tool stringer -type Platform

type Platform int

const (
	Invalid Platform = iota
	Steam
	PS3
	PS4
	XBOX360
	XBOXOne
)
