package huffman

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
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
		printCompressionStats(content, codes, bitString, compressionRatio)
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

func SerializeData(content string, verbose bool) ([]byte, error) {
	if len(content) == 0 {
		return nil, errors.New("empty content")
	}

	root := MakeTree(content)
	if root == nil {
		return nil, errors.New("failed to create tree")
	}

	serializedRoot, err := SerializeTree(root)
	if err != nil {
		return nil, err
	}

	serializedData, bitLength := Encode(root, content, verbose)

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, uint32(len(serializedRoot))); err != nil {
		return nil, err
	}

	if _, err := buf.Write(serializedRoot); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, uint32(bitLength)); err != nil {
		return nil, err
	}

	if _, err := buf.Write(serializedData); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DeserializeData(data []byte, verbose bool) (string, error) {
	if len(data) < 8 {
		return "", errors.New("invalid data length")
	}

	buf := bytes.NewReader(data)

	var rootLength uint32
	if err := binary.Read(buf, binary.LittleEndian, &rootLength); err != nil {
		return "", err
	}

	if uint32(len(data)) < 8+rootLength {
		return "", errors.New("insufficient data")
	}

	serializedRoot := make([]byte, rootLength)
	if _, err := buf.Read(serializedRoot); err != nil {
		return "", err
	}

	root, err := DeserializeTree(serializedRoot)
	if err != nil {
		return "", err
	}

	var bitLength uint32
	if err := binary.Read(buf, binary.LittleEndian, &bitLength); err != nil {
		return "", err
	}

	serializedData := make([]byte, buf.Len())
	if _, err := buf.Read(serializedData); err != nil {
		return "", err
	}

	content := Decode(root, serializedData, int(bitLength), verbose)

	return content, nil
}
