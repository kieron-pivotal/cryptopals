package operations

import (
	"github.com/kieron-pivotal/cryptopals/freqanal"
)

func SingleCharXorDecrypt(in []byte) (clear string, xorByte byte, score float64) {
	score = 1e20

	for b := byte(0); b < byte(127); b++ {
		xorBytes := []byte{b}
		xored := Xor(in, xorBytes)
		sc := freqanal.FreqScoreEnglish(string(xored))
		if sc < score {
			score = sc
			clear = string(xored)
			xorByte = b
		}
	}

	return clear, xorByte, score
}