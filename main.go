package main

import (
	"fmt"
	"os"
)

func makeTree(content string) *Node {
	fmt.Println("Original content:")
	fmt.Println(content)

	Freq := getFrequency(content)
	fmt.Println("\nCharacter frequencies:")
	for _, pair := range Freq {
		fmt.Printf("Character: '%c', Frequency: %d\n", pair.Char, pair.Count)
	}

	return buildHuffmanTree(Freq)

}

func encode(root *Node, content string) string {
	codes := make(map[rune]string)
	generateHuffmanCodes(root, "", codes)

	fmt.Println("\nHuffman Codes:")
	for char, code := range codes {
		fmt.Printf("Character: '%c', Code: %s\n", char, code)
	}

	encoded := encodeString(content, codes)

	fmt.Println("\nEncoded string:")
	fmt.Println(encoded)

	fmt.Printf("\nOriginal size: %d bits\n", len(content)*8)
	fmt.Printf("Compressed size: %d bits\n", len(encoded))
	compressionRatio := float64(len(encoded)) / float64(len(content)*8) * 100
	fmt.Printf("Compression ratio: %.2f%%\n\n", compressionRatio)
	return encoded

}

func decode(root *Node, encodedData string) string {
	decodedData := decodeHuffman(root, encodedData, 0)
	fmt.Printf("Decoded string: %s\n", decodedData)
	return decodedData
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "filename")
		fmt.Println("No file specified")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	HuffmanTreeRoot := makeTree(string(data))
	encodedData := encode(HuffmanTreeRoot, string(data))
	_ = decode(HuffmanTreeRoot, encodedData)

}
