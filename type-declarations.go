package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpYusuf NoKTP = "3277001199228380"
	var marriedStatus Married = false
	fmt.Println(noKtpYusuf)
	fmt.Println(marriedStatus)
}
