package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person
	p1.name = "黄凌波"
	p1.city = "上海"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	p2 := &person{
		name: "小王子",
		city: "北京",
		age:  20,
	}
	fmt.Printf("p2=%#v\n", p2)
}
