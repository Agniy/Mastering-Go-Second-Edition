package main

import (
	"fmt"
	"time"
)

// function is a function that prints 0-9
func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go function()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
			time.Sleep(10 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()
}
