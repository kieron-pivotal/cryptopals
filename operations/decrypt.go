package operations

import (
	"crypto/aes"

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

func RepeatingXorDecrypt(in []byte) (clear, key string) {
	probableKeyLengths := ProbableKeyLengths(in)
	minScore := 1e20
	probKey := []byte{}

	for _, l := range probableKeyLengths {

		engScore := float64(0)
		key := []byte{}
		for _, s := range SliceBytes(in, l) {
			_, x, sc := SingleCharXorDecrypt(s)
			engScore += sc
			key = append(key, x)
		}

		if engScore < minScore {
			minScore = engScore
			probKey = key
		}

	}
	clear = string(Xor(in, probKey))
	return clear, string(probKey)
}

func AES128ECBDecode(in []byte, key []byte) (clear []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	for i := 0; i*aes.BlockSize < len(in); i++ {
		block.Decrypt(in[i*aes.BlockSize:(i+1)*aes.BlockSize], in[i*aes.BlockSize:(i+1)*aes.BlockSize])
	}
	return in, nil
}
