package rlp

import (
	"testing"
)

func TestToUint(t *testing.T) {
	tests := []struct {
		input    uint64
		wantVal  uint64
	}{
		{1<<8 - 1, 1<<8 - 1},
		{1<<16 - 1, 1<<16 - 1},
		{1<<32 - 1, 1<<32 - 1},
		{1<<64 - 1, 1<<64 - 1},
	}
	for i, test := range tests {
		rlpItem := NewRLPStringFromUint(test.input)
		toUint := rlpItem.ToUint()
		if toUint != test.wantVal {
			t.Errorf("test %d: toUint mismatch, got %d want %d\ninput: %d", i, toUint, test.wantVal, test.input)
		}
	}
}