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

func gotoDemo() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

breakTag:
	fmt.Println("结束")
}
// 乘法表
func multiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {
	switchDemo()
	gotoDemo()
	
}