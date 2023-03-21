package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // send
	<-ch    // receive

	ch <- 2
	fmt.Println(<-ch)
}
