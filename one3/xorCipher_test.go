package one3

import (
	"fmt"
	"testing"
)

func TestSuccess(t *testing.T) {
	cipher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	value := Decipher(cipher)

	fmt.Println("This is the value I'm receiving: ", value)
}
