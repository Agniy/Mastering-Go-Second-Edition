package main

import (
	"fmt"
	"time"
)

func writeToChannel2(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}

func main() {
	c := make(chan int)
	go writeToChannel2(c, 10)
	time.Sleep(1 * time.Second)
}
