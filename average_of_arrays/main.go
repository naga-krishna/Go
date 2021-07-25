package main

import "fmt"

func main() {
	a := [5]int{3, 4, 12, 23, 45}
	l := len(a)
	var avg int
	for _, i := range a {
		avg += i
	}
	fmt.Println("Average of Array: ", avg/l)
}
