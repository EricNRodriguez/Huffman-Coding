package huffman

import (
	"Huffman/minHeap"
	"Huffman/bitStack"
	"bytes"
	"errors"
	"fmt"
)

type Node struct {
	Value rune
	Left  *Node
	Right *Node
}

type HuffmanTree struct {
	Root *Node
	Size int
}

func EncodeMessage(message string) (s bitStack.BitStack, huffmanTree *HuffmanTree, err error) {
	chars := []rune(message)
	charFreq := getCharacterFrequencies(chars)
	huffmanTree, err = generateHuffmanTree(charFreq)
	if err != nil {
		err = errors.New(fmt.Sprintf("[ERR] Failed huffman tree generation | %s", err.Error()))
		return
	}
	encoding := make(map[rune][]int)
	traverseAndRecord(huffmanTree.Root, encoding, []int{})
	for i := len(chars) - 1; i >= 0; i-- {
		for b := len(encoding[chars[i]]) - 1; b >= 0; b-- {
			s.Push(uint64(encoding[chars[i]][b]))
		}
	}
	return
}

func DecodeMessage(s bitStack.BitStack, huffmanTree *HuffmanTree) (message string, err error) {
	if huffmanTree == nil {
		err = errors.New("[ERR] nil huffman tree ")
		return
	}
	messageBuffer := bytes.Buffer{}
	currentNode := huffmanTree.Root
	var bit uint64
	for s.Height > 0 {
		bit, _ = s.Pop()
		if currentNode.Left == nil && currentNode.Right == nil {
			messageBuffer.WriteString(string(currentNode.Value))
			currentNode = huffmanTree.Root
		} else if bit == 1 {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	message = messageBuffer.String()
	return
}

func generateHuffmanTree(frequencies map[rune]int) (tree *HuffmanTree, err error) {
	charFreqHeap := minHeap.MinHeap{}
	for char, freq := range frequencies {
		charFreqHeap.Insert(&Node{Value: char}, freq)
	}
	size := 0
	for charFreqHeap.Size > 1 {
		minOne, err := charFreqHeap.RemoveMin()
		if err != nil {
			return nil, err
		}
		minTwo, err := charFreqHeap.RemoveMin()
		if err != nil {
			return nil, err
		}
		charFreqHeap.Insert(
			&Node{
				0,
				minOne.Value.(*Node),
				minTwo.Value.(*Node),
			},
			minOne.Frequency+minTwo.Frequency)
		size++
	}
	root, err := charFreqHeap.RemoveMin()
	if err != nil {
		return
	}
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

//tree traversal, records path to every node
func traverseAndRecord(n *Node, encoding map[rune][]int, path []int) {
	if n.Left == nil && n.Right == nil {
		encoding[n.Value] = make([]int, len(path)+1)
		copy(encoding[n.Value], path)
	}
	if n.Left != nil {
		traverseAndRecord(n.Left, encoding, append(path, 1))
	}
	if n.Right != nil {
		traverseAndRecord(n.Right, encoding, append(path, 0))
	}
	return
}
