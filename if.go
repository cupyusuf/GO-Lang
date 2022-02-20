package main

import "fmt"

func main() {
	name := "Gita"

	if name == "Yusuf" {
		fmt.Println("Hallo Yusuf")
	} else if name == "Supriadi" {
		fmt.Println("Hallo Supriadi")
	} else {
		fmt.Println("Hi, Kenalan Dong  ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
