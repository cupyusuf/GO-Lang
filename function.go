package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
	fmt.Println("Good Bye World")
	fmt.Println("Hello World")
}

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
