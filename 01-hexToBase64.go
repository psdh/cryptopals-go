package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func hexToBase64(hexInput string) string {
	rawBytes, error := hex.DecodeString(hexInput)
	if nil != error {
		log.Fatal(error)
	}
	answer := base64.StdEncoding.EncodeToString(rawBytes)

	return answer
}

func main() {
	args := os.Args[1:]

	if 1 != len(args) {
		fmt.Println("Expected exactly 1 hex string as an argument")
	}

	fmt.Println(hexToBase64(args[0]))
}
