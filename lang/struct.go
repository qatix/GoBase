package main

import (
	"fmt"
	"strings"
	"math"
)

func printBoard(s [][]string)  {
	for i:=0;i<len(s) ;i++  {
		fmt.Printf("%s\n",strings.Join(s[i]," "))
	}
}

func printSlice(s string,x []int)  {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s,len(x),cap(x),x)
}

func compute(fn func(float64,float64) float64) float64  {
	return fn(3,4)
}

func main()  {
	fmt.Println("stuct test")

	var p *int
	i := 22
	p = &i

	fmt.Println("*p:",*p)

	*p = 24
	fmt.Println("i:",i)

	//go without point op,like in c: *(p+1)

	type Vertex struct {
		x int
		y int
		z int
	}

	fmt.Println(Vertex{1,2,3})

	v := Vertex{3,4,5}
	fmt.Println(v)
	v.y  = 33
	fmt.Println(v)

	p2 := &v
	p2.z = 99
	fmt.Println(*p2)

	//array
	var a [2]string
	a[0] = "tang"
	a[1] = "xxx"
	//a[2] = "ccc"

	fmt.Println(a)

	s := []int{1,2,88,44,43,1111}
	fmt.Println("s==",s)

	for i:=0;i<len(s) ;i++  {
		fmt.Printf("s[%d] == %d\n",i,s[i])
	}

	fmt.Println(s[1:3])

	game := [][]string{
		[]string{"-","_","_"},
		[]string{"-","_","_"},
		[]string{"-","_","_"},
	}
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"
	fmt.Println(game)

	printBoard(game)

	//slice
	b := make([]int,5)
	b = b[:cap(b)]
	b = b[1:]
	fmt.Println(b)
	//    b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	//    b = b[:cap(b)] // len(b)=5, cap(b)=5
	//    b = b[1:] // len(b)=4, cap(b)=4


	var  z []int
	fmt.Println(z,len(z),cap(z))
	if z == nil{
		fmt.Println("z is nil")
	}

	z = append(z,11)
	printSlice("z",z)
	z = append(z, 22)
	printSlice("z",z)


	arr_a := [2]string{"hell","ffkfk"}
	arr_b := [...]string{"ccc","ddd","zzz"}
	fmt.Println(arr_a)
	fmt.Println(arr_b)

	bs := []byte{'a','c','z'}
	fmt.Println(bs)

	a1 := []string{"John", "Paul"}
	b1 := []string{"George", "Ringo", "Pete"}
	a1 = append(a1, b1...) // equivalent to "append(a, b[0], b[1], b[2])"
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
	fmt.Println(a1)


	//range
	var pow = []int{4,5,6,99,12}
	for i,v := range pow{
		fmt.Printf("2**%d = %d\n",i,v)
	}

	pow2 := make([]int,10)
	for i:= range pow2{
		pow2[i] = 1<<uint(i)
	}
	for _,value := range pow2{
		fmt.Printf("%d\n",value)
	}

	//map
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell La"] = Vertex{1,2,3}
	m["CCC"] = Vertex{3,4,111}
	fmt.Println(m["Bell La"])
	fmt.Println(m)

	var m2 = map[string]Vertex{
		"aa" : Vertex{1,3,0},
		"bb" : Vertex{7,3,4,},
	}
	fmt.Println(m2)
	delete(m2,"aa")
	fmt.Println(m2)

	v,ok := m2["cc"]
	fmt.Println("the value : ",v,"present?",ok)


	//func
	hypot := func(x,y float64) float64{
		return math.Sqrt(x*x+y*y)
	}
	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))

}
