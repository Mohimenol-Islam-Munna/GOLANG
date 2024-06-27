package main

import "fmt"

func main() {
	var name string
	fmt.Println("Taking user input")
	fmt.Scanln(&name)
	fmt.Println("user name : ", name)

	var age int
	fmt.Scanf("%d", &age)
	fmt.Println("user age : ", age)
}
