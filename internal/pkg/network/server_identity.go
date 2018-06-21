package network

func GetIdentity(id uint16, stype uint16) uint32 {
	rsl := uint32(stype)
	rsl = rsl << 16
	rsl = rsl + uint32(id)
	return rsl
}

func GetServerType(identity uint32) uint16 {
	tmp := identity >> 16
	return uint16(tmp)
}
