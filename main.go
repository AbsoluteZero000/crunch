package main

import (
	"flag"
	"fmt"
	"os"
)

func makeTree(content string) *Node {
	//	fmt.Println("Original content:")
	//	fmt.Println(content)

	Freq := getFrequency(content)
	// fmt.Println("\nCharacter frequencies:")
	// for _, pair := range Freq {
	// 	fmt.Printf("Character: '%c', Frequency: %d\n", pair.Char, pair.Count)
	// }

	return buildHuffmanTree(Freq)

}

func verbose_encode(root *Node, content string) string {
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

func verbose_decode(root *Node, encodedData string) string {
	decodedData := decodeHuffman(root, encodedData, 0)
	fmt.Printf("Decoded string: %s\n", decodedData)
	return decodedData
}

func encode(root *Node, content string) string {

	codes := make(map[rune]string)
	generateHuffmanCodes(root, "", codes)

	encoded := encodeString(content, codes)
	return encoded
}

func decode(root *Node, encodedData string) string {
	decodedData := decodeHuffman(root, encodedData, 0)
	return decodedData
}

func main() {
	inputFile := flag.String("i", "", "Input file name")
	outputFile := flag.String("o", "", "Output file name")
	verbose := flag.Bool("v", false, "Enable verbose mode")

	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var encodedData string
	HuffmanTreeRoot := makeTree(string(data))

	if *verbose {
		encodedData = verbose_encode(HuffmanTreeRoot, string(data))
		_ = verbose_decode(HuffmanTreeRoot, encodedData)
	} else {
		encodedData = encode(HuffmanTreeRoot, string(data))
		_ = decode(HuffmanTreeRoot, encodedData)
	}

	if *outputFile != "" {
		err = os.WriteFile(*outputFile, []byte(encodedData), 0644)
		if err != nil {
			fmt.Println("Couldn't write to file.")
			fmt.Println(err)
		}
	}

	os.Exit(0)

}
