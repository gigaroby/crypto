package ecb

import (
	"bytes"
	"errors"
)

const aesBlockLength = 16

func getBlock(line []byte, blockNumber int) []byte {
	return line[blockNumber*aesBlockLength : (blockNumber+1)*aesBlockLength]
}

func countSameBlocks(line []byte) int {
	c := 0
	n := len(line) / aesBlockLength
	if len(line)%aesBlockLength != 0 {
		panic("line length is not a multiple of aesBlockSize")
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if bytes.Equal(getBlock(line, i), getBlock(line, j)) {
				c++
			}
		}
	}
	return c
}

func DetectAESWithECB(candidates [][]byte) (int, int, error) {
	if len(candidates) < 1 {
		return 0, 0, errors.New("at least one candidate is required")
	}
	matchIdx := 0
	matchLen := countSameBlocks(candidates[0])

	for idx, candidate := range candidates[1:] {
		ml := countSameBlocks(candidate)
		if ml >= matchLen {
			matchLen = ml
			matchIdx = idx + 1
		}
	}

	return matchIdx, matchLen, nil
}
