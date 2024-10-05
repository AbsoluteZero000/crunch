package huffman

import (
	"github.com/absolutezero000/encoding/huffman"
	"testing"
)

func TestHuffmanEncodingDecoding(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "Simple test",
			input: "Hello, World!",
		},
		{
			name:  "Longer test",
			input: "Hello from NeoVim this is a long test to figure out if my huffman encoding function is working properly or not",
		},
		{
			name:  "Repeated characters",
			input: "aaaabbbcccdddeeee",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := huffman.MakeTree(tc.input)

			encoded, bitLength := huffman.Encode(root, tc.input, false)

			decoded := huffman.Decode(root, encoded, bitLength, false)

			if decoded != tc.input {
				t.Errorf("Encoding/decoding mismatch. Original: %s, Got: %s", tc.input, decoded)
			}

			if len(encoded) >= len(tc.input)*8 {
				t.Errorf("Encoding did not compress the data. Original: %d bits, Encoded: %d bits", len(tc.input)*8, len(encoded))
			}
		})
	}
}
