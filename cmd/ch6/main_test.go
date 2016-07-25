package main

import "testing"

func TestSplitBlock(t *testing.T) {
	var (
		example1 = []byte{1, 2, 3, 4, 5, 6, 7}
		example2 = []byte{1, 2, 3, 4, 5, 6}
	)

	t1 := splitBlocks(example1, 3)
	if len(t1) != 3 || len(t1[2]) != 1 {
		t.Error("error block splitting array")
	}
	t2 := splitBlocks(example2, 2)
	if len(t2) != 3 {
		t.Error("error block splitting array")
	}
}
