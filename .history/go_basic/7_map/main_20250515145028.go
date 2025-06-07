package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 顺序遍历map
func sequenceTraversal() {
	rand.Seed(time.Now().UnixNano())
	
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, scoreMap[key])
	}
	fmt.Println()
}

// 统计字符串中每个单词出现的次数


func main() {
	sequenceTraversal()
}