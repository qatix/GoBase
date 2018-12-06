package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

func add(x int,y int) int  {
	return x+y
}

func multiply(x,y int) int {
	return x*y;
}

func swap(x,y string) (string,string)  {
	return y,x
}

func split(sum int) (x,y int)  {
	x = sum*4/9
	y = sum -x
	return
}

var c,python,java bool

const (
	Big = 1<<100
	Small = Big >> 99
)

func needInt(x int) int  {
	return x*10+1
}

func needFloat(x float64) float64  {
	return x*0.1;
}

func main()  {
	fmt.Println("hello world")
	fmt.Println("sqrt:5 %g",math.Sqrt(5))
	fmt.Println("rand num:",rand.Int())
	fmt.Println("PI:",math.Pi)
	fmt.Println("Phi",math.Phi)

	fmt.Println("add:",add(4,55))
	fmt.Println("multiply:",multiply(4,8))

	a,b := swap("hello","world")
	fmt.Println("swap res",a,b)

	fmt.Println(split(27))


	var i int
	java := true
	fmt.Println(i,c,python,java)

	var j,k int = 1,2
	var o,p = true,"str-test"
	fmt.Println(j,k)
	fmt.Println(o,p)

	var(
		ToBe bool = false
		MaxInt uint64 = 1 <<64 - 1
		z complex128 = cmplx.Sqrt(-5 + 12i)
	)

	const f= "%T(%v)\n"
	fmt.Printf(f,ToBe,ToBe)
	fmt.Printf(f,MaxInt,MaxInt)
	fmt.Printf(f,z,z)

	//type transfer
	var it = 42
	var ft float64 = float64(it)
	var ut uint = uint(ft)
	fmt.Println(it,ft,ut)

	//zero value:
	//变量在定义时没有明确的初始化时会赋值为 零值
	//bool : false
	//string : ""
	//int,float : 0


	g := 0.877 + 5i
	fmt.Println(g)

	v := 44
	fmt.Printf("v is of type %T\n",v)

	//const
	//can be string,float,int,bool,charactor
	const world = "你好"
	const Truth  = true
	fmt.Println(world)
	fmt.Println(Truth)


	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big)) overflows
	//fmt.Println(Big) overflows
}
