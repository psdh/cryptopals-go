package one2

import (
	"encoding/hex"
	"log"
)

func fixedXor(buffer1, buffer2 []byte) []byte {
	if len(buffer1) != len(buffer2) {
		log.Fatal("Uhh, buffer lengths must be same ra")
	}

	var answer []byte
	answer = make([]byte, len(buffer1))

	for i := range buffer1 {
		answer[i] = buffer1[i] ^ buffer2[i]
	}

	return answer
}

func decodeHex(buffer string) []byte {
	rawBytes, error := hex.DecodeString(buffer)
	if nil != error {
		log.Fatal(error)
	}

	return rawBytes
}

// FixedXorHexEncoded Function to return Hex encoded XOR value for two hex encoded string inputs
func FixedXorHexEncoded(buf1, buf2 string) string {
	buffer1 := decodeHex(buf1)
	buffer2 := decodeHex(buf2)

	xordBuffer := fixedXor(buffer1, buffer2)

	return hex.EncodeToString(xordBuffer)
}
