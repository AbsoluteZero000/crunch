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

	HuffmanTreeRoot := huffman.MakeTree(string(data))
	encodedData, _ := huffman.Encode(HuffmanTreeRoot, string(data), *verbose)

	if *outputFile != "" {
		if err := os.WriteFile(*outputFile, encodedData, 0600); err != nil {
			log.Printf("Error writing to output file %s: %v", *outputFile, err)
		}
	}

	// treeData, _ := huffman.SerializeTree(HuffmanTreeRoot)
	//
	// fmt.Println(treeData)
	//
	// treeRoot, _ := huffman.DeserializeTree(treeData)
	// _ = huffman.Decode(treeRoot, encodedData, bitLength, *verbose)
}
