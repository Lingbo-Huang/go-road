package main

import "fmt"

func switchDemo() {
	s := "a"
	switch {
		case s == "a":
			fmt.Println("a")
			fallthrough
		case s == "b":
			fmt.Println("b")
		case s == "c":
			fmt.Println("c")
	}
}

func main() {
	
}