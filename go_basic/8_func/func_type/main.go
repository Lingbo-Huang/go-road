package main

import (
	"errors"
	"fmt"
)

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	var c calculation
	c = add
	fmt.Printf("type of c: %T\n", c)
	fmt.Println(c(1, 2))

	f := add
	fmt.Printf("type of f: %T\n", f)
	fmt.Println(c(1, 2))

	ret2 := calc(10, 20, add)
	fmt.Println(ret2)
}
