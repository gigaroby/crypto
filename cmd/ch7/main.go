package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gigaroby/crypto/english"
	"github.com/gigaroby/crypto/xor"
)

var (
	input = flag.String("input", "data.txt", "file to load")
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Open(*input)
	handleErr(err)
	defer f.Close()
	i := byte(0)
	det := &english.Detector{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		h, err := hex.DecodeString(scanner.Text())
		handleErr(err)
		for i = 0; i < 255; i++ {
			det.Add(xor.SingleByteXOR(h, i))
		}
	}

	w, s := det.Best()
	fmt.Printf("best: %s, score: %f\n", string(w), s)
}
