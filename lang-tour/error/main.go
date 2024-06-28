package main

import "fmt"

func main() {
	fmt.Println("Hello Error")
	var money float64 = 500
	var count float64 = 0

	amount := money / count

	fmt.Println("amount:", amount)

	fmt.Println("next fmt:")
}
