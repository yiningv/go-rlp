package rlp

const (
	OffsetShortString byte = 0x80
	OffsetLongString  byte = 0xb7 // 0x80 + 0x37(55)
	OffsetShortList   byte = 0xc0
	OffsetLongList    byte = 0xf7 // 0xc0 + 0x37(55)
)
