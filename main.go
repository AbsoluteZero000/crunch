package main

import (
	"flag"
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
	var encodedData string
	HuffmanTreeRoot := makeTree(string(data))

	if *verbose {
		encodedData = verbose_encode(HuffmanTreeRoot, string(data))
		_ = verbose_decode(HuffmanTreeRoot, encodedData)
	} else {
		encodedData = encode(HuffmanTreeRoot, string(data))
	}

	if *outputFile != "" {
		outputData := seralizeTreeAndEncodedData(HuffmanTreeRoot, encodedData)
		if err := os.WriteFile(*outputFile, []byte(outputData), 0600); err != nil {
			log.Printf("Error writing to output file %s: %v", *outputFile, err)
		}
	}

}
