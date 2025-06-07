package main

import "fmt"

func switchDemo() {
	s := "a"
	switch {
		case s == "a":
			fmt.Println("a")
			fallthrough
		case s == "b":
	}
}

func main() {
	
}