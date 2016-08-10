package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"log"
	"os"

	"github.com/gigaroby/crypto/ecb"
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
	sc := bufio.NewScanner(f)

	data := make([][]byte, 0)
	for sc.Scan() {
		d := make([]byte, hex.DecodedLen(len(sc.Bytes())))
		hex.Decode(d, sc.Bytes())
		data = append(data, d)
	}
	handleErr(sc.Err())

	idx, n, err := ecb.DetectAESWithECB(data)
	handleErr(err)

	log.Printf("best idx: %d, n: %d\n", idx, n)
}
