package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	Char  rune
	Freq  int
	Left  *Node
	Right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Freq < pq[j].Freq }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Pair struct {
	Char  rune
	Count int
}

func getFrequency(s string) []Pair {
	frequencyMap := make(map[rune]int)
	for _, char := range s {
		frequencyMap[char]++
	}
	pairs := make([]Pair, 0, len(frequencyMap))
	for char, count := range frequencyMap {
		pairs = append(pairs, Pair{Char: char, Count: count})
	}

	return pairs
}

func buildHuffmanTree(pairs []Pair) *Node {
	pq := make(PriorityQueue, len(pairs))
	for i, pair := range pairs {
		pq[i] = &Node{Char: pair.Char, Freq: pair.Count}
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)
		newNode := &Node{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}
		heap.Push(&pq, newNode)
	}

	return heap.Pop(&pq).(*Node)
}

func generateHuffmanCodes(root *Node, prefix string, codes map[rune]string) {
	if root == nil {
		return
	}

	if root.Char != 0 {
		codes[root.Char] = prefix
		return
	}

	generateHuffmanCodes(root.Left, prefix+"0", codes)
	generateHuffmanCodes(root.Right, prefix+"1", codes)
}

func decodeHelper(root *Node, encoded string, index int) (decodedChar string, newIndex int) {
	if root.Char != 0 {
		return string(root.Char), index
	}

	if encoded[index] == '0' {
		return decodeHelper(root.Left, encoded, index+1)
	} else {
		return decodeHelper(root.Right, encoded, index+1)
	}

}

func decodeHuffman(root *Node, encoded string, index int) string {
	decodedData := ""
	char := ""
	for index < len(encoded) {
		char, index = decodeHelper(root, encoded, index)
		decodedData += char
	}

	return decodedData
}

func encodeString(s string, codes map[rune]string) string {
	var encoded string
	for _, char := range s {
		encoded += codes[char]
	}
	return encoded
}

func makeTree(content string) *Node {
	Freq := getFrequency(content)
	return buildHuffmanTree(Freq)
}

func encode(root *Node, content string) (data []byte, bitlength int) {
	codes := make(map[rune]string)
	generateHuffmanCodes(root, "", codes)

	var bitString string
	for _, char := range content {
		bitString += codes[char]
	}

	return bitStringToBytes(bitString), len(bitString)
}

func decode(root *Node, encodedData []byte, bitLength int) string {
	bitString := bytesToBitString(encodedData, bitLength)

	return decodeHuffman(root, bitString, 0)
}

func verbose_encode(root *Node, content string) (data []byte, bitLength int) {
	fmt.Println("Original content:")
	fmt.Println(content)

	codes := make(map[rune]string)
	generateHuffmanCodes(root, "", codes)

	fmt.Println("\nHuffman Codes:")
	for char, code := range codes {
		fmt.Printf("Character: '%c', Code: %s\n", char, code)
	}

	bitString := encodeString(content, codes)

	fmt.Println("\nBit string before conversion:")
	fmt.Println(bitString)

	encoded := bitStringToBytes(bitString)

	fmt.Printf("\nOriginal size: %d bits\n", len(content)*8)
	fmt.Printf("Compressed size: %d bits\n", len(bitString))
	compressionRatio := float64(len(bitString)) / float64(len(content)*8) * 100
	fmt.Printf("Compression ratio: %.2f%%\n\n", compressionRatio)

	return encoded, len(bitString)
}

func verbose_decode(root *Node, encodedData []byte, bitLength int) string {
	bitString := bytesToBitString(encodedData, bitLength)

	fmt.Println("Recovered bit string:")
	fmt.Println(bitString)
	fmt.Println()

	decodedData := decodeHuffman(root, bitString, 0)
	fmt.Printf("Decoded string: %s\n", decodedData)
	return decodedData
}
