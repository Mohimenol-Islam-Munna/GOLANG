package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		recoverPanic := recover()
		if recoverPanic != nil {
			fmt.Println("Panic happened. Please take proper action. : ", recoverPanic)
		}

		fmt.Println("go defer 1 after everything. everything is done")
	}()
	//defer func() {
	//	time.Sleep(5 * time.Second)
	//	fmt.Println("go defer 2 after everything")
	//}()
	//defer func() {
	//	fmt.Println("go defer 3 after everything")
	//}()

	fmt.Println("after defer")
	time.Sleep(5 * time.Second)
	fmt.Println("after sleep")

	var age int
	fmt.Print("Please enter your age: ")
	fmt.Scanln(&age)

	fmt.Println("We are checking your age. please wait.")
	time.Sleep(5 * time.Second)

	if age < 18 {
		fmt.Println("age is less than 18")
		panic("You are in under age")
	} else {
		fmt.Println("You are able to do further work. let's go forward.")
	}

	fmt.Println("after age check")
}
