package one3

import (
	"bytes"
	"testing"
)

func TestSuccess(t *testing.T) {
	cipher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decipheredBytes := Decipher(cipher)

	expectedDecipheredBytes := []byte{99, 79, 79, 75, 73, 78, 71, 0, 109, 99, 7, 83, 0, 76, 73, 75, 69, 0, 65, 0, 80, 79, 85, 78, 68, 0, 79, 70, 0, 66, 65, 67, 79, 78}

	if !bytes.Equal(expectedDecipheredBytes, decipheredBytes) {
		t.Errorf("That doesn't seem right. Expected: %v, Got: %v", expectedDecipheredBytes, decipheredBytes)
	}
}
