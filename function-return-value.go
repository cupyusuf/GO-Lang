package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hallo Bro"
	} else {
		return "Hallo " + name
	}
}

func main() {
	result := getHello("Yusuf")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
