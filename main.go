package main

import (
	"Huffman/huffman"
	"Huffman/minHeap"
	"fmt"
)

func main() {
	minHeap := minHeap.MinHeap{}
	fmt.Println("here")
	minHeap.Insert("-5", -5)

	minHeap.Insert(2, 2)
	minHeap.Insert(1, 1)
	minHeap.Insert(3, 3)
	minHeap.Insert(2, 2)
	minHeap.Insert(2, 2)

	minHeap.Insert(5, 5)
	minHeap.Insert(4, 4)
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())
	//fmt.Println(minHeap.Remove())

	freq, _, err := huffman.EncodeMessage("abcde aa bbb c")
	fmt.Println(freq, err)


}
