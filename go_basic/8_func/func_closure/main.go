package main

import (
	"fmt"
	"strings"
)

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	var f = adder(5)
	fmt.Println(f(10)) //15
	fmt.Println(f(20)) // 35
	fmt.Println(f(30)) // 65

	f1 := adder(10)
	fmt.Println(f1(20)) //30
	fmt.Println(f1(50)) //80

	fmt.Println(f(10)) //75

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test"))

}
