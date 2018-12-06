package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum //send sum into c
}

func main() {
	a := []int{1, 3, 4, 5, -1, 7, 8}
	c := make(chan int, 100)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}
