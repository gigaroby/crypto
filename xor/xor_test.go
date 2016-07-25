package xor

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestFromHexString(t *testing.T) {
	source := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b, err := hex.DecodeString(source)
	if err != nil {
		t.Errorf("error converting hex: %s", err)
	}
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if got := base64.StdEncoding.EncodeToString(b); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestFixedXor(t *testing.T) {
	s1 := "1c0111001f010100061a024b53535009181c"
	s2 := "686974207468652062756c6c277320657965"
	h1, e1 := hex.DecodeString(s1)
	h2, e2 := hex.DecodeString(s2)
	if e1 != nil || e2 != nil {
		t.Error("invalid char in input")
	}

	x, err := FixedXOR(h1, h2)
	if err != nil {
		t.Errorf("error XORing strings: %s", err)
	}

	expected := "746865206b696420646f6e277420706c6179"
	if got := hex.EncodeToString(x); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestRepeatingKeyXOR(t *testing.T) {
	data := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	result := RepeatingKeyXOR(data, key)
	if got := hex.EncodeToString(result); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
