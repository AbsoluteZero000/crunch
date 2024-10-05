package huffman

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func MakeTree(content string) *Node {
	Freq := getFrequency(content)
	return buildHuffmanTree(Freq)
}

func Encode(root *Node, content string, verbose bool) (data []byte, bitLength int) {

	codes := make(map[rune]string)
	generateHuffmanCodes(root, "", codes)

	bitString := encodeString(content, codes)
	encoded := bitStringToBytes(bitString)
	compressionRatio := float64(len(bitString)) / float64(len(content)*8) * 100

	if verbose {
		fmt.Println("Original content:")
		fmt.Println(content)

		fmt.Println("\nHuffman Codes:")
		for char, code := range codes {
			fmt.Printf("Character: '%c', Code: %s\n", char, code)
		}

		fmt.Println("\nBit string before conversion:")
		fmt.Println(bitString)

		fmt.Printf("\nOriginal size: %d bits\n", len(content)*8)
		fmt.Printf("Compressed size: %d bits\n", len(bitString))
		fmt.Printf("Compression ratio: %.2f%%\n\n", compressionRatio)
	}

	return encoded, len(bitString)
}

func Decode(root *Node, encodedData []byte, bitLength int, verbose bool) string {
	bitString := bytesToBitString(encodedData, bitLength)
	decodedData := decodeHuffman(root, bitString, 0)

	if verbose {
		fmt.Println("Recovered bit string:")
		fmt.Println(bitString)
		fmt.Println()
		fmt.Printf("Decoded string: %s\n", decodedData)
	}
	return decodedData
}

func SerializeTree(root *Node) ([]byte, error) {
	if root == nil {
		return nil, nil
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(root)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DeserializeTree(data []byte) (*Node, error) {
	if data == nil {
		return nil, nil
	}

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var root Node
	err := dec.Decode(&root)
	if err != nil {
		return nil, err
	}

	return &root, nil
}
