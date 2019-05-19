package main

import (
	"Huffman/minHeap"
	"fmt"
)

func main() {
	minHeap := minHeap.MinHeap{}
	minHeap.Insert('2', 2)
	minHeap.Insert(rune(10), -10)
	minHeap.Insert('5',5)
	minHeap.Insert('7',7)

	fmt.Println(minHeap.Remove())
	fmt.Println(minHeap.Remove())
	fmt.Println(minHeap.Remove())
	fmt.Println(minHeap.Remove())
	fmt.Println(minHeap.Remove())
	fmt.Println(minHeap.Remove())


}