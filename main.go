package main

import (
	"flag"
	"github.com/absolutezero000/encoding/huffman"
	"log"
	"os"
)

func main() {
	inputFile := flag.String("i", "", "Input file name")
	outputFile := flag.String("o", "", "Output file name")
	verbose := flag.Bool("v", false, "Enable verbose mode")

	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", *inputFile, err)
	}
	var encodedData []byte
	var bitLength int
	HuffmanTreeRoot := huffman.MakeTree(string(data))

	encodedData, bitLength = huffman.Encode(HuffmanTreeRoot, string(data), *verbose)
	_ = huffman.Decode(HuffmanTreeRoot, encodedData, bitLength, *verbose)

	if *outputFile != "" {
		if err := os.WriteFile(*outputFile, encodedData, 0600); err != nil {
			log.Printf("Error writing to output file %s: %v", *outputFile, err)
		}
	}

}
