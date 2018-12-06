package main

import "fmt"

/*
#include <stdlib.h>
*/
import "C" //这一行必须在*/注释的下一行

func main() {
	fmt.Println("Hello World!")
	fmt.Println(C.rand())
}
