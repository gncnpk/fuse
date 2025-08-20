package fobcolor

type EFobColor uint8

//go:generate go tool stringer -type=EFobColor
const (
	Plain1TerracottaOrange EFobColor = iota // #B77440
	Plain2SlateBlue                         // #556574
	Plain3DarkSlateGray                     // #4E5255
	Plain4ChestnutBrown                     // #7C483B
	Plain5LightSlateGray                    // #918B8B
	Plain6OliveDrab                         // #6A6747
	Plain7DustyRose                         // #B77674
	Plain8Mocha                             // #97866A

	Camo1Square
	Camo2Desert
	Camo3TigerStripe
	Camo4Veil
	Camo5Woodland
	Camo6MuddyBrown
	Camo7Christmas
	Camo8Forest // 0xf
	Camo9Zebra
	Camo10Wooden
	Camo11Rust
	Camo12YellowWashedOut
	End
)

func IsValid(color int) bool {
	return color > -1 && color < int(End)
}
