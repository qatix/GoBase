package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	fmt.Println("push 1")

	ch <- 4
	fmt.Println("push 4")

	ch <- 8
	//fatal error: all goroutines are asleep - deadlock!
	fmt.Println("push 8")

}
