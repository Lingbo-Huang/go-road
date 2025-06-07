package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 顺序遍历map
func sequenceTraversal() {
	rand.Seed(time.Now().UnixNano())
	
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		
	}
}

func main() {

}