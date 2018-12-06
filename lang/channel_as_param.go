package main

import "fmt"

func main()  {

	messages := make(chan string)

	go work(messages)

	msg := <- messages
	fmt.Println("main func received:",msg)
}

func work(messages chan<- string)  {
	messages <- "msg from work"
}
