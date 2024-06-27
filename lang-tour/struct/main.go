package main

import "fmt"

func main() {
	type Person struct {
		name    string
		age     int
		country string
	}
	var player Person
	player = Person{"messi", 35, "argentina"}
	fmt.Println("player :", player.name)

	var playerPtr *Person
	playerPtr = &player
	fmt.Println("Player Pointer variable :", playerPtr)
	fmt.Println("player ptr :", (*playerPtr).name)

	//update
	player.name = "Ronaldo"
	fmt.Println("player :", player)
}
