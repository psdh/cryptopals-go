package one3

import (
	"encoding/hex"
	"log"
	"sort"
	"strings"
)

type byFrequency []string

const numberOfChars = 26
const frequentString1 string = "ETAOIN"
const frequentString2 string = "SHRDLU"

func decodeHex(buffer string) []byte {
	rawBytes, error := hex.DecodeString(buffer)
	if nil != error {
		log.Fatal(error)
	}

	return rawBytes
}

func frequentCharacterCount(str string) int {
	answer := 0
	for _, c := range str {
		if strings.ContainsRune(frequentString1, c) {
			answer += 2
		} else if strings.ContainsRune(frequentString2, c) {
			answer++
		}
	}
	return answer
}

func (s byFrequency) Len() int {
	return len(s)
}

func (s byFrequency) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byFrequency) Less(i, j int) bool {
	return frequentCharacterCount(s[i]) < frequentCharacterCount(s[j])
}

// Decipher  This will decipher them all...
func Decipher(cipher string) string {
	cipherBytes := decodeHex(cipher)

	var possibleAnswers []string
	possibleAnswers = make([]string, numberOfChars)
	for i := 0; i < numberOfChars; i++ {
		a := byte('a' + i)
		answer := make([]byte, len(cipherBytes))

		for i, b := range cipherBytes {
			answer[i] = a ^ b
		}

		possibleAnswers[i] = strings.ToUpper(string(answer))
	}

	sort.Sort(byFrequency(possibleAnswers))

	return possibleAnswers[len(possibleAnswers)-1]
}
