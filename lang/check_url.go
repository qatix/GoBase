package main

import (
	"net/url"
	"fmt"
)

func isValidUrl(urlstr string)  bool {
	_,err := url.ParseRequestURI(urlstr)
	if err != nil{
		return false
	}else{
		return true
	}
}

func main()  {

	// = true
	fmt.Println(isValidUrl("http://www.golangcode.com"))

	// = false
	fmt.Println(isValidUrl("http://golangcode.com?#abc"))

	// = false
	fmt.Println(isValidUrl("cccc"))
}
