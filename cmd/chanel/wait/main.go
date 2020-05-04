package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for c := range ch  {
		fmt.Printf("%d \n", c)
	}
}
