package main

import "fmt"

type shape interface {
	area() float64
	price() float64
}

type house struct {
	location  string
	length    float64
	breadth   float64
	unitPrice float64
}

func main() {
	fmt.Println("interface main function")
	house1 := house{"dhaka", 12.9, 34.9, 12.8}
	fmt.Println("house1", house1)

	area := house1.calculateArea()
	fmt.Println("area", area)
}

func (h house) calculateArea() float64 {
	fmt.Println("h : ", h)
	area := h.length * h.breadth
	return area
}

func (h house) area() float64 {
	return h.length * h.breadth
}

func (h house) price() float64 {
	return h.length * h.breadth * h.unitPrice
}
