package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var CLOSEA = false

var DATA = make(map[uint64]bool)

func random(min, max uint64) uint64 {
	return uint64(rand.Int63n(int64(max-min))) + uint64(min)
}

func first(min, max uint64, out chan<- uint64) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- uint64, in <-chan uint64) {
	for x := range in {
		//fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan uint64) {
	var sum uint64
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random2 numbers is %d.\n", sum)
	fmt.Println("Len of DATA is: ", len(DATA))
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two uint64eger parameters!")
		return
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n11 := uint64(n1)
	n2, _ := strconv.Atoi(os.Args[2])
	n22 := uint64(n2)

	if n1 > n2 {
		fmt.Println("%d should be smaller than %d.\n", n11, n22)
		return
	}

	A := make(chan uint64)
	B := make(chan uint64)

	go first(n11, n22, A)
	go second(B, A)
	third(B)
}
