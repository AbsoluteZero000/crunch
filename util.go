package main

import (
	"bytes"
	"math"
)

func bitStringToBytes(s string) []byte {
	numBytes := int(math.Ceil(float64(len(s)) / 8.0))
	result := make([]byte, numBytes)

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			byteIndex := i / 8
			bitIndex := 7 - (i % 8)
			result[byteIndex] |= 1 << bitIndex
		}
	}

	return result
}

func bytesToBitString(data []byte, bitLength int) string {
	var result bytes.Buffer

	for i := 0; i < bitLength; i++ {
		byteIndex := i / 8
		bitIndex := 7 - (i % 8)

		if data[byteIndex]&(1<<bitIndex) != 0 {
			result.WriteByte('1')
		} else {
			result.WriteByte('0')
		}
	}

	return result.String()
}
