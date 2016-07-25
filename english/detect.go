package english

import (
	"bytes"
	"math"
)

var (
	freq = map[byte]float64{
		' ': 0.1918182,
		'e': 0.1041442,
		't': 0.0729357,
		'a': 0.0651738,
		'o': 0.0596302,
		'n': 0.0564513,
		'i': 0.0558094,
		's': 0.0515760,
		'r': 0.0497563,
		'h': 0.0492888,
		'd': 0.0349835,
		'l': 0.0331490,
		'u': 0.0225134,
		'c': 0.0217339,
		'm': 0.0202124,
		'f': 0.0197881,
		'w': 0.0171272,
		'g': 0.0158610,
		'y': 0.0145984,
		'p': 0.0137645,
		'b': 0.0124248,
		'v': 0.0082903,
		'k': 0.0050529,
		'x': 0.0013692,
		'j': 0.0009033,
		'q': 0.0008606,
		'z': 0.0007836,
	}
)

func ScoreWord(w []byte) float64 {
	var (
		score   = float64(0)
		relFreq = make(map[byte]float64)
	)
	w = bytes.ToLower(w)
	for _, c := range w {
		if _, ok := freq[c]; !ok {
			continue
		}
		relFreq[c] += 1
	}

	for k, v := range freq {
		vv, ok := relFreq[k]
		if !ok {
			score += v
			continue
		}
		vv = vv / float64(len(w))
		score += math.Abs(v - vv)
	}

	return score
}

type Detector struct {
	match      []byte
	matchValue float64
}

func (d *Detector) Add(word []byte) {
	v := ScoreWord(word)
	if d.match == nil || d.matchValue > v {
		d.match = word
		d.matchValue = v
	}
}

func (d *Detector) Best() (word []byte, score float64) {
	return d.match, d.matchValue
}
