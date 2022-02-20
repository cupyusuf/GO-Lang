package main

import "fmt"

func getFullName() (string, string, string) {
	return "Yusuf", "Supriadi", "Golang"
}

func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}
