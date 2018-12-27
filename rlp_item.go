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
	encodeBytes := encodeString(s.value)
	return encodeBytes
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
	encodeBytes := encodeList(l.items)
	return encodeBytes
}

func (l *RLPList) AddItem(items ...RLPItem)  {
	l.items = append(l.items, items...)
}