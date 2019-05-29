package bitArray

import (
	"fmt"
	"bytes"
	"strconv"
	"math"
)

type IntegerSet []uint64

func (s *IntegerSet) Add(x uint64) {
	cell, bit := math.Floor(float64(x)/64), x%64
	for uint64(cell) >= uint64(len(*s)) {
		*s = append(*s, 0)
	}
	(*s)[uint64(cell)] |= (1 << bit)
}
func (s *IntegerSet) Remove(x uint64) {
	cell, bit := x/64, x%64
	if cell < uint64(len(*s)) {
		(*s)[cell] ^= (1 << bit)
	}
}
func (s *IntegerSet) String() string {
	var buffer bytes.Buffer
	for i := 0; i < len(*s)*64; i++ {
		cell, bit := i/64, i%64
		if ((*s)[cell] & (1 << uint64(bit))) != 0 {
			buffer.WriteString(fmt.Sprintf("%s ",strconv.Itoa(i)))
		}
	}
	return buffer.String()
}
func (s *IntegerSet) Len() int {
	c := 0
	for i := 0; i < len(*s)*64; i++ {
		cell, bit := i/64, i%64
		if ((*s)[cell] & (1 << uint64(bit))) != 0 {
			c += 1
		}
	}
	return c
}
func (s *IntegerSet) Clear() {
	for i := 0; i < len(*s); i++ {
		(*s)[i] = 0
	}
}
func (s *IntegerSet) Union(setTwo *IntegerSet) {
	for i := 0; i < len(*setTwo); i++ {
		if i < len(*s) {
			(*s)[i] |= (*setTwo)[i]
		} else {
			*s = append(*s, (*setTwo)[i])
		}
	}
}