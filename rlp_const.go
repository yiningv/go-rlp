package rlp

const (
	OFFSET_SHORT_STRING byte = 0x80
	OFFSET_LONG_STRING  byte = 0xb7 // 0x80 + 0x37(55)
	OFFSET_SHORT_LIST   byte = 0xc0
	OFFSET_LONG_LIST    byte = 0xf7 // 0xc0 + 0x37(55)
)
