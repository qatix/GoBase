package main

import (
	"strconv"
	"fmt"
)

func main()  {

	var myNumber uint64
	myNumber = 18446744073709551615

	str := strconv.FormatUint(myNumber,10)

	fmt.Println("The number is:" + str)

}
