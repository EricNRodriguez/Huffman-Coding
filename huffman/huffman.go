package huffman

import (
	"Huffman/minHeap"
	"fmt"
)

func EncodeMessage(message string) (encodedMessage []rune, encoding map[rune]int, err error) {
	frequencies := getCharacterFrequencies([]rune(message))
	encoding, err = encodeCharacters(frequencies)

	return
}

func encodeCharacters(charFreq map[rune]int) (encoding map[rune]int, err error) {
	charFreqHeap := minHeap.MinHeap{}
	for char, freq := range charFreq {
		charFreqHeap.Insert(char, freq)
	}
	for charFreqHeap.Size > 1 {
		minOne, _ := charFreqHeap.RemoveMin()
		minTwo, _ := charFreqHeap.RemoveMin()
		value := []interface{}{
			nil,
			minOne,
			minTwo,
		}
		charFreqHeap.Insert(value, minOne.Frequency + minTwo.Frequency)
	}
	huffmanTree, _ := charFreqHeap.RemoveMin()

	fmt.Println(huffmanTree.Value)

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