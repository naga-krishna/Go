package main

import "fmt"

func main() {
	r := 0.0
	const Pi = 3.14
	fmt.Print("Enter Radius of the circle: ")
	fmt.Scan(&r)
	fmt.Println("Area of the circle: ", Pi*r*r)
	fmt.Println("Circumference of the circle: ", 2*Pi*r)
}
