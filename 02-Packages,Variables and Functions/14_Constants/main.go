package main

import "fmt"

const Pi = 3.14

func main() {
	const Hi = "你好"
	fmt.Print(Hi, "World")
	fmt.Println("Happy", Pi, "Day")
	const Truth = "真"
	fmt.Println(Truth, "or", "假")
}
