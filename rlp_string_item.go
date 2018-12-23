package rlp

// A string (ie. byte array) is an item
type RLPStringItem struct {
	value []byte
}

func NewStringItem(value []byte) *RLPStringItem {
	return &RLPStringItem{value:value}
}


