package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Yusuf"
	sayHelloTo(firstName, "Supriadi")
	sayHelloTo("Shalma", "Shabila")
}