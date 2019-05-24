package huffman

import (
	"Huffman/minHeap"
)

type Node struct {
	Value rune
	Left *Node
	Right *Node
}

type HuffmanTree struct {
	Root *Node
	Size int
}

//TODO:
// - Return encoded Message as io reader
// - Write Decoder

func EncodeMessage(message string) (encodedMessage []int, huffmanTree *HuffmanTree, err error) {
	charFreq := getCharacterFrequencies([]rune(message))
	huffTree, err := buildHuffmanTree(charFreq)
	encoding := make(map[rune][]int)
	traverseAndRecord(huffTree.Root, encoding, []int{})
	for _, char := range []rune(message) {
		for _, b := range encoding[char] {
			encodedMessage = append(encodedMessage, b)
		}
	}
	return
}

func buildHuffmanTree(frequencies map[rune]int) (tree *HuffmanTree, err error) {
	charFreqHeap := minHeap.MinHeap{}
	for char, freq := range frequencies {
		charFreqHeap.Insert(&Node{Value:char}, freq)
	}
	size := 0
	for charFreqHeap.Size > 1 {
		minOne, _ := charFreqHeap.RemoveMin()
		minTwo, _ := charFreqHeap.RemoveMin()
		charFreqHeap.Insert(
			&Node{
				69,
				minOne.Value.(*Node),
				minTwo.Value.(*Node),
			},
			minOne.Frequency + minTwo.Frequency)
		size ++
	}
	root, _ := charFreqHeap.RemoveMin()
	tree = &HuffmanTree{
		root.Value.(*Node),
		size,
	}
	return
}

func getCharacterFrequencies(chars []rune) (frequencies map[rune]int) {
	frequencies = make(map[rune]int)
	for _, char := range chars {
		if _, ok := frequencies[char]; ok {
			frequencies[char] += 1
		} else {
			frequencies[char] = 1
		}
	}
	return
}

//in order traversal, records path to every node
func traverseAndRecord(n *Node, encoding map[rune][]int, path []int) {
	if n.Left == nil && n.Right == nil {
		encoding[n.Value] = path
	}
	if n.Left != nil {
		traverseAndRecord(n.Left, encoding, append(path, 1))
	}
	if n.Right != nil {
		traverseAndRecord(n.Right, encoding, append(path, 0))
	}
	return
}


