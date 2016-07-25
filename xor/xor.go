package xor

import (
	"errors"
	"math"

	"github.com/gigaroby/crypto/english"
)

func FixedXOR(h1, h2 []byte) ([]byte, error) {
	if len(h1) != len(h2) {
		return nil, errors.New("inputs have different length")
	}

	res := make([]byte, len(h1))
	for i := 0; i < len(h1); i++ {
		res[i] = h1[i] ^ h2[i]
	}
	return res, nil
}

func SingleByteXOR(h1 []byte, c byte) []byte {
	res := make([]byte, len(h1))
	for i, b := range h1 {
		res[i] = b ^ c
	}
	return res
}

func RepeatingKeyXOR(data, key []byte) []byte {
	out := make([]byte, len(data))
	for i, d := range data {
		out[i] = d ^ key[i%len(key)]
	}
	return out
}

func GuessSingleByteXORKey(data []byte) byte {
	var (
		key  byte
		best float64 = math.MaxFloat64
	)

	for k := byte(0); k < 255; k++ {
		perm := SingleByteXOR(data, k)
		score := english.ScoreWord(perm)
		if score < best {
			best = score
			key = k
		}
	}

	return key
}
