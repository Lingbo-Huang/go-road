package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s: %v len(s): %d cap(s): %d\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)，而不是len(s)
	fmt.Printf("s2: %v len(s2): %d cap(s2): %d\n", s2, len(s2), cap(s2))
	t := a[2:3:4] // t := a[low:high:max]
}