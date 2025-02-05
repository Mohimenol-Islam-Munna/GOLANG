package main

import (
	"fmt"
	"3.GO_WITH_TESTS/2.integers"
)

func main() {
	fmt.Println("Hello world", hello("Munna"))

	fmt.Println("sum :", integers.Add(1,2))
}

func hello(name string) string {return name}