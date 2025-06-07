package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s: %v len(s): %d cap(s): %d\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)，而不是len(s)
}