package one1

import "testing"

func TestSuccess(t *testing.T) {
	hexEncodedStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64EncodedStr := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	base64EncodedStr := HexToBase64(hexEncodedStr)

	if expectedBase64EncodedStr != base64EncodedStr {
		t.Errorf("That doesn't seem right. Expected: %s, Got: %s", expectedBase64EncodedStr, base64EncodedStr)
	}
}
