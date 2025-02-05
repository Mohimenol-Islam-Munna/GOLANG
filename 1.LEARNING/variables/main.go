package main

import "fmt"

func main() {
	fmt.Println("variables")
	// variable with var keyword
	// string
	var name string
	name = "munna"
	fmt.Println("name ::", name)
	fmt.Printf("name type :%T:\n", name)

	// type inference
	var anotherName = "ifti"
	fmt.Println("anotherName ::", anotherName)
	fmt.Printf("another Name is :%v:\n", anotherName)
	fmt.Printf("anotherName type :%T:\n", anotherName)

	// integer
	// there are variation on int type variable like , uint8, uint64, int8, int64
	var id int = 160133
	fmt.Println("student id is :", id)

	var profit float32 = 12.3746374347
	var income float64 = 12.3746374347
	fmt.Println("profit ::", profit)
	fmt.Printf("profit type :%T:\n", profit)

	fmt.Println("income ::", income)
	fmt.Printf("income type :%T:\n", income)

	// variable without var keyword
	// this type variable declaration does not work outside function
	isAdmin := true
	isStaff := false
	fmt.Println("isAdmin :", isAdmin)
	fmt.Println("isStaff :", isStaff)

	fmt.Println("isAdmin :", isAdmin)
	fmt.Printf("anotherName type :%T:\n", anotherName)
	fmt.Println("isStaff :", isStaff)
	fmt.Printf("isStaff type :%T:\n", isStaff)

	// const variable
	const campus string = "Pabna University of Science and Technology"
	fmt.Println("campus :", campus)
	fmt.Printf("campus type :%T:\n", campus)
}
