package main

import "fmt"

func twoSum(nums []int, target int) [][]int {
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func main() {
	a := [...]int{1:1, 3:5}
	fmt.Println(a)
	fmt.Printf("len: %d cap: %d\n", len(a), cap(a))
	fmt.Printf("type of a: %T\n", a)

	nums := []int{2, 7, 11, 15}
}