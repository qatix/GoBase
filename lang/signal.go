package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
)

const FILE_NAME  = "go-example.txt"

func main()  {
	SetupCloseHander();

	CreateFile()
	for{
		fmt.Println("- Sleeping")
		time.Sleep(10*time.Second)
	}
}

func SetupCloseHander()  {
	c := make(chan os.Signal,2)
	signal.Notify(c,os.Interrupt,syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		DeleteFile()
		os.Exit(0)
	}()
}

func DeleteFile()  {
	fmt.Println("- Run Clean Up - Delete example file")
	_ = os.Remove(FILE_NAME)
	fmt.Println("- Good bye!")
}

func CreateFile()  {
	fmt.Println("- Create example file")
	file,_ := os.Create(FILE_NAME)
	defer file.Close()
}
