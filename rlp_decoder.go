package rlp

import "encoding/binary"

func decode(b []byte) RLPItem {
	// 根据前缀信息决定处理类型
	// len(b) == 0 -> 0x80 or 0xc0
	if len(b) == 0 {

	}
	// [0x00, 0x7f]
	if b[0] < OFFSET_SHORT_STRING {
		item := NewRLPString(b)
		return item
	}
	// [0x80, 0xb7]
	if b[0] >= OFFSET_SHORT_STRING && b[0] <= OFFSET_LONG_STRING {
		strLen := int(b[0] - 0x80)
		if len(b) < 1+strLen {
			panic("invalid input")
		}
		item := NewRLPString(b[1 : 1+strLen])
		return item
	}

	// [0xb8, 0xbf]
	if b[0] > OFFSET_LONG_STRING && b[0] < OFFSET_SHORT_LIST {
		strLenLen := int(b[0] - 0xB7)
		if len(b) < strLenLen+1 {
			panic("invalid input")
		}
		strLen := int(toUint(b[1 : strLenLen+1]))
		if len(b) < 1+strLenLen+strLen {
			panic("invalid input")
		}
		item := NewRLPString(b[1+strLenLen : 1+strLenLen+strLen])
		return item
	}

	// [0xc0, 0xf7]
	if b[0] >= OFFSET_SHORT_LIST && b[0] <= OFFSET_LONG_LIST {
		listStr := int(b[0] - 0xC0)
		if len(b) < 1+listStr {
			panic("invalid input")
		}
		item := NewRLPList(decode(b[1:]))
		return item
	}

	// [0xf8, 0xff]
	if b[0] > OFFSET_LONG_LIST {
		listLenLen := int(b[0] - 0xF7)
		if len(b) < listLenLen+1 {
			panic("invalid input")
		}
		listLen := int(toUint(b[1 : listLenLen+1]))
		if len(b) < 1+listLenLen+listLen {
			panic("invalid input")
		}
		item := NewRLPList(decode(b[listLenLen:]))
		return item
	}

	return nil
}

func toUint(v []byte) uint64 {
	b := make([]byte, 8)
	copy(b[8-len(v):], v)
	return binary.BigEndian.Uint64(b)
}
