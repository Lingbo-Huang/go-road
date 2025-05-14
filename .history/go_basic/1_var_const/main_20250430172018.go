package main

import "fmt"

func foo() (int, string){
	return 10, "hlb"
}

// 变量和常量
func main() {
	var a int 
	a = 10
	var s string = "hello"
	var name, age = "hlb", 18
	n := 10
	m := 200
	x, _ := foo()
	_, y := foo()
	fmt.Println(x, y)
	// var a int = 10
	fmt.Println(a, s)
	fmt.Printf("%s 已经 %d 岁了\n", name, age)
	fmt.Println(n + m)
}