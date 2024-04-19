package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	server := NewAPIServer(":3000")
	server.Run()
}
