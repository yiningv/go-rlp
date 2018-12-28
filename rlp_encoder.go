package rlp

import (
	"encoding/binary"
)

// for string: 0xb7-0x80=0x37 55
// for list:   0xf7-0xc0=0x37 55
func encode(b []byte, offset byte) []byte {
	if len(b) <= 55 {
		result := make([]byte, len(b)+1)
		result[0] = offset + byte(len(b))
		copy(result[1:], b)
		return result
	} else {
		lenBytes := trimLeftZerosFromBytes(uintToBytes(uint64(len(b))))
		lenBytesLen := len(lenBytes)
		result := make([]byte, len(b)+lenBytesLen+1)
		result[0] = offset + 0x37 + byte(lenBytesLen)
		copy(result[1:lenBytesLen+1], lenBytes)
		copy(result[lenBytesLen+1:], b)
		return result
	}
}

// 转换uint为字节数组
func uintToBytes(strLen uint64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, strLen)
	return result[:]
}

// 去掉字节数组的前导零
func trimLeftZerosFromBytes(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		if b[i] != 0 {
			result := make([]byte, len(b)-i)
			copy(result, b[i:])
			return result
		}
	}
	return []byte{}
}
