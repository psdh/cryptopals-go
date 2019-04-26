package one2

import "testing"

func TestSuccess(t *testing.T) {
	hexBuf1 := "1c0111001f010100061a024b53535009181c"
	hexBuf2 := "686974207468652062756c6c277320657965"

	expectedHexXordBuffer := "746865206b696420646f6e277420706c6179"
	hexXordBuffer := FixedXorHexEncoded(hexBuf1, hexBuf2)

	if expectedHexXordBuffer != hexXordBuffer {
		t.Errorf("Bad Bad Bad \n Expected: %s, Got: %s", expectedHexXordBuffer, hexXordBuffer)
	}
}
