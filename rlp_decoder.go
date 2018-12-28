package rlp

import "encoding/binary"

func Decode(b []byte) RlpItem {
	return decode(b)[0]
}

func decode(b []byte) []RlpItem {
	if len(b) == 0 {
		panic("invalid input")
	}
	// [0x00, 0x7f]
	if b[0] < OffsetShortString {
		var items []RlpItem
		item := NewRlpString(b[0:1])
		items = append(items, item)
		if len(b) > 1 {
			items = append(items, decode(b[1:])...)
		}
		return items
	}
	// [0x80, 0xb7]
	if b[0] >= OffsetShortString && b[0] <= OffsetLongString {
		strLen := int(b[0] - 0x80)
		if len(b) < 1+strLen {
			panic("invalid input")
		}
		var items []RlpItem
		item := NewRlpString(b[1 : 1+strLen])
		items = append(items, item)
		if len(b) > 1+strLen {
			items = append(items, decode(b[1+strLen:])...)
		}
		return items
	}

	// [0xb8, 0xbf]
	if b[0] > OffsetLongString && b[0] < OffsetShortList {
		strLenLen := int(b[0] - 0xB7)
		if len(b) < strLenLen+1 {
			panic("invalid input")
		}
		strLen := int(toUint(b[1 : strLenLen+1]))
		if len(b) < 1+strLenLen+strLen {
			panic("invalid input")
		}
		var items []RlpItem
		item := NewRlpString(b[1+strLenLen : 1+strLenLen+strLen])
		items = append(items, item)
		if len(b) > 1+strLenLen+strLen {
			items = append(items, decode(b[1+strLenLen+strLen:])...)
		}
		return items
	}

	// [0xc0, 0xf7]
	if b[0] >= OffsetShortList && b[0] <= OffsetLongList {
		listStr := int(b[0] - 0xC0)
		if len(b) < 1+listStr {
			panic("invalid input")
		}
		item := NewRlpList(decode(b[1:])...)
		items := []RlpItem{item}
		return items
	}

	// [0xf8, 0xff]
	if b[0] > OffsetLongList {
		listLenLen := int(b[0] - 0xF7)
		if len(b) < listLenLen+1 {
			panic("invalid input")
		}
		listLen := int(toUint(b[1 : listLenLen+1]))
		if len(b) < 1+listLenLen+listLen {
			panic("invalid input")
		}
		item := NewRlpList(decode(b[listLenLen:])...)
		items := []RlpItem{item}
		return items
	}

	return nil
}

func toUint(v []byte) uint64 {
	b := make([]byte, 8)
	copy(b[8-len(v):], v)
	return binary.BigEndian.Uint64(b)
}
