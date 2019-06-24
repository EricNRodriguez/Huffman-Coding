package test

import (
	"Huffman/bitStack"
	"math/rand"
	"testing"
	"time"
)

func TestRandomStack(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(100)
	bits := []uint64{}
	stack := bitStack.BitStack{}
	var bit uint64
	for i := 0; i < length; i++ {
		bit = uint64(rand.Intn(126) % 2)
		stack.Push(bit)
	}
	index := len(bits) - 1
	for i := len(bits) - 1; i >= 0; i-- {
		bit, err := stack.Pop()
		if err != nil || bit != bits[index] {
			t.Fail()
		}
		index--
	}
	return
}
