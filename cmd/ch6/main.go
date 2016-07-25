package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/gigaroby/crypto/hamming"
	"github.com/gigaroby/crypto/xor"
)

const (
	MinKeySize = 2
	MaxKeySize = 40
)

var (
	input = flag.String("input", "data.txt", "file to load")
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func splitBlocks(content []byte, blockSize int) [][]byte {
	blocks, reminder := len(content)/blockSize, len(content)%blockSize
	res := make([][]byte, blocks)
	for i := 0; i < blocks; i++ {
		j := i * blockSize
		res[i] = content[j : j+blockSize]
	}

	if reminder > 0 {
		res = append(res, content[len(content)-reminder:])
	}
	return res
}

func transpose(content [][]byte) [][]byte {
	ret := make([][]byte, len(content[0]))
	for _, c := range content {
		for j := 0; j < len(content[0]); j++ {
			// safeguard, last array may be shorter than first
			if j >= len(c) {
				break
			}
			ret[j] = append(ret[j], c[j])
		}
	}

	return ret
}

func averageDistance(parts [][]byte) float64 {
	p := len(parts)
	if len(parts[p-1]) != len(parts[p-2]) {
		parts = parts[:p-1]
	}
	acc := float64(0)
	n := 0
	for i := 0; i < len(parts)-1; i++ {
		for j := i + 1; j < len(parts); j++ {
			d, err := hamming.EditDistance(parts[i], parts[j])
			handleErr(err)
			acc += float64(d)
			n++
		}
	}
	return acc / float64(n)
}

func bestKeyLength(content []byte) int {
	distance := math.MaxFloat64
	best := 0
	for keyLen := MinKeySize; keyLen < MaxKeySize+1; keyLen++ {
		parts := splitBlocks(content, keyLen)
		d := averageDistance(parts) / float64(keyLen)
		if d < distance {
			distance = d
			best = keyLen
		}
	}
	return best
}

func main() {
	f, err := os.Open(*input)
	handleErr(err)
	defer f.Close()
	decoder := base64.NewDecoder(base64.StdEncoding, f)
	content, err := ioutil.ReadAll(decoder)
	handleErr(err)

	keyLen := bestKeyLength(content)
	parts := splitBlocks(content, keyLen)
	transpos := transpose(parts)
	finalKey := make([]byte, len(transpos))
	for i, t := range transpos {
		finalKey[i] = xor.GuessSingleByteXORKey(t)
	}

	plain := xor.RepeatingKeyXOR(content, finalKey)
	log.Println(string(plain))
}
