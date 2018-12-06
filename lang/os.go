package main

import (
	"fmt"
	"runtime"
	"os"
)

func main() {
	fmt.Println("os:", runtime.GOOS)
	fmt.Println(os.Getgid())
	fmt.Println(os.Hostname())
}
