package english

import (
	"encoding/hex"
	"testing"

	"github.com/gigaroby/crypto/xor"
)

func TestDetector(t *testing.T) {
	input, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		t.Errorf("error converting input: %s", err)
	}

	det := &Detector{}
	i := byte(0)
	for i = 0; i < 255; i++ {
		det.Add(xor.SingleByteXOR(input, i))
	}

	expected := "Cooking MC's like a pound of bacon"
	match, _ := det.Best()
	if string(match) != expected {
		t.Errorf("expected %s, got %s", expected, string(match))
	}
}
