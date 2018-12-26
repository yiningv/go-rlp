package rlp

import (
	"testing"
)

func TestEncodeString(t *testing.T) {
	fromUint := NewRLPItemFromUint(12111111111111111111)
	bytes := encodeString(fromUint.value)
	t.Log(bytes)
}