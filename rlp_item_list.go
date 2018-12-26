package rlp

import (
	"reflect"
)

// A list of items is an item
type RLPItemList struct {
	items []RLPItem
}

func NewRLPItemList(items []RLPItem) *RLPItemList {
	return &RLPItemList{items:items}
}

func (l *RLPItemList) AddItem(items ...RLPItem)  {
	l.items = append(l.items, items...)
}

func (l *RLPItemList) AddItemList(itemLists ...RLPItemList) {
}

func typeOf(v interface{}) string {
	return reflect.TypeOf(v).String()
}