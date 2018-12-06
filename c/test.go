package main

import "fmt"

/*
int add(int a,int b){
return a+b;
}
*/
import "C" //这一行必须在*/注释的下一行

func main() {
	fmt.Println("Hello World!")
	fmt.Println(C.add(2, 1))
}
