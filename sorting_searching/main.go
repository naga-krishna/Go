package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{23, 1, 7, 34, 2}
	sort.Ints(x)
	fmt.Println(x)
	fmt.Println("Location of 7:", 1+sort.SearchInts(x, 7))
	y := []string{"India", "USA", "Japan", "China", "Austria"}
	sort.Strings(y)
	fmt.Println(y)
	fmt.Println("Location of USA:", 1+sort.SearchStrings(y, "USA"))
	z := []float64{23.23, 1.1, 7.7, 3.4, 202.9}
	sort.Float64s(z)
	fmt.Println(z)
	fmt.Println("Location of 3.4:", 1+sort.SearchFloat64s(z, 3.4))
}
