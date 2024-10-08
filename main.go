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
	decode := flag.Bool("d", false, "Decode mode")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", *inputFile, err)
	}
	var outputData []byte

	if *decode {
		decodedData, err := huffman.DeserializeData(data, *verbose)
		if err != nil {
			log.Fatalf("Failed to deserialize data: %v", err)
		}
		os.Stdout.Write([]byte(decodedData))
		outputData = []byte(decodedData)
	} else {
		outputData, err = huffman.SerializeData(string(data), *verbose)
		if err != nil {
			log.Fatalf("Failed to serialize data: %v", err)
		}
	}

	if *outputFile != "" {
		if err := os.WriteFile(*outputFile, outputData, 0600); err != nil {
			log.Printf("Error writing to output file %s: %v", *outputFile, err)
		}
	}

}
