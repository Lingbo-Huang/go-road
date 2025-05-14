package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 10
	fmt.Printf("%T %v %d %b\n", a, a, a, a)

	var b int = 077
	fmt.Printf("%T %v %d %b\n", b, b, b, b)
	var c int = 0xff
	fmt.Printf("%T %v %d %b\n", c, c, c, c)

	fmt.Printf("%f\n", math.Pi)

	var c1 complex64 = 1 + 2i
	fmt.Println(c1)

	s1 := `第一行
	第二行
	第三行`
	
}