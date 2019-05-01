package one3

import (
	"encoding/hex"
	"fmt"
	"log"
	"sort"
	"strings"
)

type byFrequency [26]string

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
	str = strings.ToUpper(str)
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
	fmt.Println("comparing ", s[i][:4], " ", frequentCharacterCount(s[i]), " with ", s[j][:4], " ", frequentCharacterCount(s[j]))
	return frequentCharacterCount(s[i]) < frequentCharacterCount(s[j])
}

// Decipher  This will decipher them all...
func Decipher(cipher string) string {
	cipherBytes := decodeHex(cipher)

	var possibleAnswers [26]string
	for i := 0; i < 26; i++ {
		a := byte('a' + i)
		answer := make([]byte, len(cipherBytes))

		for i, b := range cipherBytes {
			answer[i] = a ^ b
		}

		fmt.Println(string(answer))
		possibleAnswers[i] = string(answer)
	}

	sort.Sort(byFrequency(possibleAnswers))

	fmt.Println(possibleAnswers[len(possibleAnswers)-3:])

	for _, answer := range possibleAnswers {
		fmt.Println(frequentCharacterCount(answer))
	}

	return possibleAnswers[len(possibleAnswers)-1]
}
