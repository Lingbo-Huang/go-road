package main

import "fmt"

func switchDemo() {
	s := "a"
	switch {
		case s == "a":
			fmt.Println("a")
			fallthrough // fallthrough 会继续执行下一个 case
		case s == "b":
			fmt.Println("b")
		case s == "c":
			fmt.Println("c")
		default:
			fmt.Println("default")
	}
}

func main() {
	switchDemo()
}