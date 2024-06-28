package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Print("Enter Positive Integer Number : ")
	var number int
	var _, err = fmt.Scanln(&number)

	if err != nil {
		inputError := fmt.Errorf("something went wrong. Please provide appropriate number")
		fmt.Println("inputError : ", inputError)
	}

	fmt.Print("Enter Positive Float Number : ")
	var number2 float64
	var _, _ = fmt.Scanln(&number2)

	if number2 < 18.0 {
		fmt.Println("Number is too small")
		number2Error := errors.New("number is too small")
		fmt.Println("number2Error : ", number2Error)
	} else {
		fmt.Println("Number is ok.")
	}
}
