package one3

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	cipher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decipheredText := Decipher(cipher)

	expectedDecipheredText := "COOKINGMCSLIKEAPOUNDOFBACON"

	if expectedDecipheredText != decipheredText {
		t.Errorf("That doesn't seem right. Expected: %s, Got: %s", expectedDecipheredText, decipheredText)
	}
}
