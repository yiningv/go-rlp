package rlp

import (
	"encoding/binary"
	"fmt"
)

type RLPItem interface {
	EncodeRLP() []byte
}

// A string (ie. byte array) is an item
type RLPString struct {
	value []byte
}

func NewRLPString(value []byte) *RLPString {
	return &RLPString{value: value}
}

func NewRLPStringFromStr(str string) *RLPString {
	return NewRLPString([]byte(str))
}

func NewRLPStringFromUint(uintVal uint64) *RLPString {
	valBytes := uintToBytes(uintVal)

	return NewRLPString(trimLeftZerosFromBytes(valBytes))
}

func (s *RLPString) EncodeRLP() []byte {
	b := s.value
	if len(b) == 0 {
		return []byte{0x80}
	}
	if len(b) == 1 && b[0] < 0x7F {
		return b
	}
	return encode(b, OFFSET_SHORT_STRING)
}

func (s *RLPString) GetBytes() []byte {
	return s.value
}

func (s *RLPString) ToString() string {
	return fmt.Sprintf("%s", s.value)
}

func (s *RLPString) ToUint() uint64 {
	b := make([]byte, 8)
	copy(b[8-len(s.value):], s.value)
	return binary.BigEndian.Uint64(b)
}

// A list of items is an item
type RLPList struct {
	items []RLPItem
}

func NewRLPList(items ...RLPItem) *RLPList {
	return &RLPList{items:items}
}

func (l *RLPList) EncodeRLP() []byte {
	if len(l.items) == 0 {
		return []byte{0xC0}
	}
	var b []byte
	for _, item := range l.items {
		b = append(b, item.EncodeRLP()...)
	}
	return encode(b, OFFSET_SHORT_LIST)
}

func (l *RLPList) AddItem(items ...RLPItem)  {
	l.items = append(l.items, items...)
}