package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g > %g\n", v, lim)
	}
	return lim
}

func main() {

	fmt.Println("condition learn")

	//go only exist for loop
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("sum(1-99):", sum)

	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Println("sum2:", sum2)

	sum3 := 1
	for sum3 < 2000 {
		sum3 += sum3
	}
	fmt.Println("sum3:", sum3)

	//for  { // infinite loop
	//fmt.Println("aaaa")
	//}

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//swith
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	//time
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	switch  time.Saturday {
	case today+0:
		fmt.Println("Today")
	case today+1:
		fmt.Println("Tomorrow")
	case today+2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	tnow := time.Now()
	switch  {
	case tnow.Hour() < 12:
		fmt.Println("Good morning")
	case tnow.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}


	//defer
	defer fmt.Println("first print")
	fmt.Println("second print")

	fmt.Println("Counting")
	for i:=0;i<10;i++{
		defer fmt.Println("count i:",i)
	}
	fmt.Println("Counting done")
}
