package clustersecurityparam

import (
	"fmt"
	"github.com/unknown321/fuse/guardrank"
	"github.com/unknown321/fuse/weaponrange"
)

// see tpp::net::FobTarget::ReceiveDetailInfo

// ClusterSecurityParam is set in `security settings -> advanced -> all decks`
type ClusterSecurityParam struct {
	DefenseLevel   byte // security rating, shield icon, 1-127
	NonLethal      byte // 1 for nonlethal
	SwimsuitID     byte // 0-36
	GuardRank      guardrank.EGuardRank
	EquipmentGrade byte // 1-11
	WeaponRange    weaponrange.EWeaponRange
	HasGuards      byte // 0-1
	Bit25_31       byte // bits 25-31 (7 bits), always 0?
}

// FUN_1407ef170, 0x1407ef3d2
func (c *ClusterSecurityParam) FromInt(i int) error {
	c.DefenseLevel = byte((i >> 0xC) & 0x7f)
	c.NonLethal = byte((i >> 0xA) & 0x1)
	c.SwimsuitID = byte((i >> 0x13) & 0x3f)
	gr := byte((i >> 0x6) & 0xf)
	if !guardrank.IsValid(gr) {
		return fmt.Errorf("invalid guard rank %d", gr)
	}
	c.GuardRank = guardrank.EGuardRank(gr)
	c.EquipmentGrade = byte((i >> 2) & 0xf)
	if c.EquipmentGrade < 1 {
		c.EquipmentGrade = 1
	}
	if c.EquipmentGrade > 0xf {
		c.EquipmentGrade = 0xf
	}

	wr := byte(i & 0x3)
	if !weaponrange.IsValid(wr) {
		return fmt.Errorf("invalid weapon range %d", wr)
	}

	c.WeaponRange = weaponrange.EWeaponRange(wr)

	c.HasGuards = byte((i >> 0xB) & 0x1)  // bit 11
	c.Bit25_31 = byte((i >> 0x19) & 0x7f) // bits 25-31

	return nil
}

func (c *ClusterSecurityParam) ToInt() int {
	securityRatingPart := int(c.DefenseLevel&0x7f) << 0xC
	nonLethal := int(c.NonLethal&0x1) << 0xA
	swimsuit := int(c.SwimsuitID&0x3f) << 0x13
	guardRank := int(c.GuardRank&0xf) << 0x6
	eqGrade := int(c.EquipmentGrade&0xf) << 0x2
	weaponRange := int(c.WeaponRange & 0x3)

	guards := int(c.HasGuards&0x1) << 0xB
	b25 := int(c.Bit25_31&0x7f) << 0x19

	return securityRatingPart | nonLethal | swimsuit |
		guardRank | eqGrade | weaponRange | guards | b25
}
