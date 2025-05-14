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

	fmt.Println(s1)

	s2 := "hello小鹿"
	for _, r := range s2 {
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	changeString()
}

func changeString() {
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "黄萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
	sqrtDemo()
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func countChineseCha