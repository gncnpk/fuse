package constructparam

import (
	"fmt"
	"github.com/unknown321/fuse/areacode"
	"github.com/unknown321/fuse/fobcolor"
)

// see tpp::net::FobTarget::ReceiveDetailInfo
// BaseSvarsToMbsParam

// /chunk0/Assets/tpp/pack/mbdvc/mb_dvc_top_fpk/Assets/tpp/level/location/mtbs/block_layout
var layouts = []int{
	0,
	1,
	2,
	3,
	10,
	11,
	12,
	13,
	20,
	21,
	22,
	23,
	30,
	31,
	32,
	33,
	40,
	41,
	42,
	43,
	53,
	60,
	61,
	62,
	63,
	70,
	71,
	72,
	73,
	80,
	81,
	82,
	83,
	90,
	91,
	92,
	93,
	100,
	101,
	102,
	103,
	500,
	900,
	950,
	952,
	953,
	971,
}

type ConstructParam struct {
	AreaCode   areacode.EAreaCode
	Color      fobcolor.EFobColor
	LayoutCode int // TODO known as areaID // see layouts list above
	Mysterious int // ?????? 0-61, most likely up to 63

	// color is 0x_BIT2_xxxx_BIT1_xx
	baseColorBit1 int // color ID, 0-7 plain, bit2 + 8-f,0-3 = camo
	baseColorBit2 int // color ID2, if set then color = bit1 + 0xf
	bit0          int // always 1
}

func (c *ConstructParam) FromInt(i int) error {
	c.baseColorBit1 = (i >> 0x8) & 0xf
	c.baseColorBit2 = (i >> 0x1c) & 0x1
	col := c.baseColorBit1
	if c.baseColorBit2 == 1 {
		col += 0xf + 1
	}

	c.Color = fobcolor.EFobColor(col)

	areaID := ((i >> 0xc) & 0xffff) & 0x3ff
	if !areacode.IsValid(areaID) {
		return fmt.Errorf("invalid area code %d", areaID)
	}
	c.AreaCode = areacode.EAreaCode(areaID)

	c.LayoutCode = (i >> 0x1) & 0x7f
	layoutValid := false
	for _, v := range layouts {
		if v == c.LayoutCode {
			layoutValid = true
			break
		}
	}

	if !layoutValid {
		return fmt.Errorf("invalid layout %d", c.LayoutCode)
	}

	c.Mysterious = (i >> 0x16) & 0x3f
	c.bit0 = i & 0x1

	return nil
}

func (c *ConstructParam) ToInt() int {
	if c.Color != fobcolor.End {
		if c.Color > 0xf {
			c.baseColorBit2 = 1
			c.baseColorBit1 = int(c.Color) - 0xf - 1
		} else {
			c.baseColorBit2 = 0
			c.baseColorBit1 = int(c.Color)
		}
	}

	layout := (c.LayoutCode & 0x7f) << 0x1
	color1 := (c.baseColorBit1 & 0xf) << 0x8
	area := (int(c.AreaCode) & 0x3ff) << 0xc
	mystery := (c.Mysterious & 0x3f) << 0x16
	color2 := (c.baseColorBit2 & 0x1) << 0x1c
	//one := c.bit0 & 0x1
	one := 1

	return layout | color1 | area | mystery | color2 | one
}
