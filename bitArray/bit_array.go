package main

import (
	"errors"
	"math"
)

type BitArray struct {
	Array 	[]uint64
	length 	int
}

func (b *BitArray) Add(x int) error {
	if x != 1 && x != 0 {
		return errors.New("Non bit element inserted into array (1 and 0)")
	}
	if b.length == len(b.Array) * 64 {
		b.Array = append(b.Array, uint64(x))
		b.length ++
		return nil
	}
	if x == 0 {
		b.Array[len(b.Array)-1] = b.Array[len(b.Array)-1] << 1
	} else {
		b.Array[len(b.Array)-1] = (b.Array[len(b.Array)-1] << 1) | 1
	}
	b.length ++
	return nil
}

func (b *BitArray) Pop() (x int, err error) {
	if b.length == 0 {
		return 0, errors.New("No elements remaining in bit array")
	}
	// reducing size of the array as elements are removed
	if int(math.Ceil(float64(b.length/64))) < len(b.Array) {
		b.Array = b.Array[:int(math.Ceil(float64(b.length/64)))+1]
	}
	if (b.Array[len(b.Array)-1] >> uint64((b.length % 64)-1)) % 2 == 0 {
		x = 0
	} else {
		x = 1
	}
	b.length--
	return
}



