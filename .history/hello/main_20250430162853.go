package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
// mac 执行 windows 下的 go 程序
// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o hello.exe main.go
// mac 执行 linux 下的 go 程序
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello main.go