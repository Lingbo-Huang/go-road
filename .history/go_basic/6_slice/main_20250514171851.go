package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s: %v len(s): %d cap(s): %d\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)，而不是len(s)
	fmt.Printf("s2: %v len(s2): %d cap(s2): %d\n", s2, len(s2), cap(s2))
	t := a[2:3:4] // t := a[low:high:max], max-low 是容量
	fmt.Printf("t: %v len(t): %d cap(t): %d\n", t, len(t), cap(t))

	b := make([]int, 2, 10)
	fmt.Printf("b: %v len(b): %d cap(b): %d\n", b, len(b), cap(b))
	c := []int{1, 2, 3}
	fmt.Printf("c: %v len(c): %d cap(c): %d\n", c, len(c), cap(c))
	// 用len(s) == 0 来判断一个切片是否为空
	// 赋值拷贝前后两个切片共享底层数组，对一个切片的修改会影响另一个切片的内容
	// copy() 函数可以将一个切片拷贝到另一个切片中

	d := []int{1, 2, 3}
	e := make([]int, 5, 5)
	copy(e, d) // copy(dst, src)
	fmt.Println(d, e)

	// 删除元素要用拼接
	f := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 删除索引为2的元素
	f = append(f[:2], f[3:]...)
	fmt.Println(f)
	
}