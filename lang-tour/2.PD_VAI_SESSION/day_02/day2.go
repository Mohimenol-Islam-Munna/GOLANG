package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println("x", x)

	fmt.Println("repeat :", Repeat("a"))

}

func Add(a int, b int ) int{
	return a + b
}

func Repeat(character string) string {
    var repeated string
    for i := 0; i < 5; i++ {
        repeated = repeated + character
    }
    return repeated
}