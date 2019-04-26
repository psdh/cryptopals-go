package one1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToBase64(hexInput string) string {
	rawBytes, error := hex.DecodeString(hexInput)
	if nil != error {
		log.Fatal(error)
	}
	answer := base64.StdEncoding.EncodeToString(rawBytes)

	return answer
}
