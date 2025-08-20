package localbaseparam

// see tpp::net::FobTarget::ReceiveDetailInfo

// TODO verify

type LocalBaseParam struct {
	PlatformsBuilt int
	Mystery0       int // ?? always 0
	Mystery1       int
}

func (c *LocalBaseParam) FromInt(i int) error {
	c.PlatformsBuilt = (i >> 0xC) & 0xf
	c.Mystery0 = (i >> 0xA) & 0x3
	c.Mystery1 = i & 0x1
	return nil
}

func (c *LocalBaseParam) ToInt() int {
	platformCountPart := (c.PlatformsBuilt & 0xf) << 0xC
	platformsBuiltPart := (c.Mystery0 & 0x3) << 0xA
	mystery1Part := c.Mystery1 & 0x1

	return platformCountPart | platformsBuiltPart | mystery1Part
}
