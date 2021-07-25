package main

import "fmt"

func main() {
	var x int64
	fmt.Print("Enter Number to get Factorial: ")
	fmt.Scan(&x)
	if x < 0 {
		fmt.Println("Factorial for negative numbers are not possbile")
		return
	}
	y := Factorial(x)
	fmt.Println(y)
}

func Factorial(x int64) int64 {
	var i, fact int64
	fact = 1
	for i = 1; i <= x; i++ {
		fact *= i
	}
	return fact
}
