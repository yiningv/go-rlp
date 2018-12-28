package rlp

import (
	"encoding/binary"
	"fmt"
)

type RlpItem interface {
	EncodeRLP() []byte
}

// A string (ie. byte array) is an item
type RlpString struct {
	value []byte
}

func NewRlpString(value []byte) *RlpString {
	return &RlpString{value: value}
}

func NewRlpStringFromStr(str string) *RlpString {
	return NewRlpString([]byte(str))
}

func NewRlpStringFromUint(uintVal uint64) *RlpString {
	valBytes := uintToBytes(uintVal)

	return NewRlpString(trimLeftZerosFromBytes(valBytes))
}

func (s *RlpString) EncodeRLP() []byte {
	b := s.value
	if len(b) == 0 {
		return []byte{0x80}
	}
	if len(b) == 1 && b[0] < 0x7F {
		return b
	}
	return encode(b, OffsetShortString)
}

func (s *RlpString) GetBytes() []byte {
	return s.value
}

func (s *RlpString) ToString() string {
	if len(s.value) == 0 {
		return ""
	}
	return fmt.Sprintf("%s", s.value)
}

func (s *RlpString) ToUint() uint64 {
	if len(s.value) == 0 {
		return 0
	}
	b := make([]byte, 8)
	copy(b[8-len(s.value):], s.value)
	return binary.BigEndian.Uint64(b)
}

// A list of items is an item
type RlpList struct {
	items []RlpItem
}

func NewRlpList(items ...RlpItem) *RlpList {
	return &RlpList{items: items}
}

func (l *RlpList) EncodeRLP() []byte {
	if len(l.items) == 0 {
		return []byte{0xC0}
	}
	var b []byte
	for _, item := range l.items {
		b = append(b, item.EncodeRLP()...)
	}
	return encode(b, OffsetShortList)
}

func (l *RlpList) AddItem(items ...RlpItem)  {
	l.items = append(l.items, items...)
}