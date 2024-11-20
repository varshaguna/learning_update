package main

import (
	"fmt"
	"math/bits"
)

func main() {
	num := uint(0b11011)
	sum := bits.OnesCount(num)
	fmt.Printf("The sum of Binary number of %b is %d", num, sum)
}
