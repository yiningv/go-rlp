package rlp

import (
	"fmt"
)

// A string (ie. byte array) is an item
// StringItem改成RLPItem, String容易让人误解只能是string数据
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

func (s *RLPItem) GetBytes() []byte {
	return s.value
}

func (s *RLPItem) ToString() string {
	return fmt.Sprintf("%s", s.value)
}
