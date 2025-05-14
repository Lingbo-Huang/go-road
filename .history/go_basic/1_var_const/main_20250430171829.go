package main

import "fmt"

// 变量和常量
func main() {
	var a int 
	a = 10
	var s string = "hello"
	var name, age = "hlb", 18
	
	// var a int = 10
	fmt.Println(a, s)
	fmt.Printf("%s 已经 %d 岁了\n", name, age)
}