package rlp

// A list of items is an item
type RLPItemList struct {
	items []RLPItem
}

func NewRLPItemList(items []RLPItem) *RLPItemList {
	return &RLPItemList{items:items}
}