package main

import (
	"bytes"
	"container/heap"
	"encoding/binary"
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

func verbose_encode(root *Node, content string) string {
	fmt.Println("Original content:")
	fmt.Println(content)

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

func seralizeTreeAndEncodedData(root *Node, encodedData string) []byte {
	// seralizedTree, err := SerializeTree(root)
	// fmt.Printf("Seralized tree: %s\n", seralizedTree)
	// if err != nil {
	// 	log.Fatalf("Failed to serialize tree: %v", err)
	// }
	// TODO: fix this and make it return the seralized tree and data
	return []byte(encodedData)

}

func SerializeTree(root *Node) ([]byte, error) {
	var buf bytes.Buffer
	err := serializeNode(&buf, root)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func serializeNode(buf *bytes.Buffer, node *Node) error {
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		err := buf.WriteByte(1)
		if err != nil {
			return err
		}

		err = binary.Write(buf, binary.LittleEndian, node.Char)
		if err != nil {
			return err
		}

		err = binary.Write(buf, binary.LittleEndian, int32(node.Freq))
		if err != nil {
			return err
		}
	} else {
		err := buf.WriteByte(0)
		if err != nil {
			return err
		}

		err = serializeNode(buf, node.Left)
		if err != nil {
			return err
		}

		err = serializeNode(buf, node.Right)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeserializeTree(data []byte) (*Node, error) {
	buf := bytes.NewReader(data)
	return deserializeNode(buf)
}

func deserializeNode(buf *bytes.Reader) (*Node, error) {
	nodeType, err := buf.ReadByte()
	if err != nil {
		return nil, err
	}

	if nodeType == 1 {
		var char rune
		err = binary.Read(buf, binary.LittleEndian, &char)
		if err != nil {
			return nil, err
		}

		var freq int32
		err = binary.Read(buf, binary.LittleEndian, &freq)
		if err != nil {
			return nil, err
		}

		return &Node{Char: char, Freq: int(freq)}, nil
	}

	left, err := deserializeNode(buf)
	if err != nil {
		return nil, err
	}

	right, err := deserializeNode(buf)
	if err != nil {
		return nil, err
	}

	return &Node{Left: left, Right: right}, nil
}
