package main

import "fmt"
/*
除了一个数字之外，其他数字都出现了两次，找到这个数字
*/

// 两个相同数字异或为0
// 0与任何数字异或为该数字本身
func findSingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {

}
