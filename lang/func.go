package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x,y float64
}


//Go 没有类。然而，仍然可以在结构体类型上定义方法。
//方法接收者 出现在 func关键字和方法名之间的参数中。
func (v *Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x+v.y*v.y)
}


//你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
//但是，不能对来自其他包的类型或基础类型定义方法。
type MyFloat float64

func (f MyFloat) abs() float64  {
	if f < 0{
		return float64(-f);
	}
	return float64(f)
}

type Abser interface {
	abs() float64
}

func main()  {

	fmt.Println("func test")
	v := &Vertex{4,5}
	fmt.Println(v.abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.abs())

	var a Abser

	a = f
	fmt.Println(a.abs())
	a = v
	fmt.Println(a.abs())



}
