package rlp

import (
	"encoding/binary"
	"fmt"
)

// A string (ie. byte array) is an item
type RLPItem struct {
	value []byte
}

func NewRLPItem(value []byte) *RLPItem {
	return &RLPItem{value: value}
}

func NewRLPItemFromStr(str string) *RLPItem {
	return NewRLPItem([]byte(str))
}

func NewRLPItemFromUint(uintVal uint64) *RLPItem {
	valBytes := uintToBytes(uintVal)

	return NewRLPItem(trimLeftZerosFromBytes(valBytes))
}

func (r *RLPItem) GetBytes() []byte {
	return r.value
}

func (r *RLPItem) ToString() string {
	return fmt.Sprintf("%s", r.value)
}

func (r *RLPItem) ToUint() uint64 {
	b := make([]byte, 8)
	copy(b[8-len(r.value):], r.value)
	return binary.BigEndian.Uint64(b)
}
