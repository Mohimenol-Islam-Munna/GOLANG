package main

import "fmt"

func main() {
	fmt.Println("Pointer")
	var name string = "munna"
	var namePtr *string = &name

	fmt.Println("name :", name)
	fmt.Println("namePtr:", namePtr)
	fmt.Println("*namePtr:", *namePtr)

	var number int = 123
	var numberAddress *int = &number

	fmt.Println("Number : ", number)
	fmt.Println("numberAddress:", numberAddress)
}
