package main

import (
	"fmt"
	"sort"
)

func main() {
	s := [5]int{12, 3, 24, 1, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(s[:])))
	fmt.Println(s)
	s1 := []string{"India", "Australia", "Korea", "USA", "Japan"}
	sort.Sort(sort.Reverse(sort.StringSlice(s1)))
	fmt.Println(s1)
	s2 := []string{"5", "2", "6", "3", "1", "4"} // unsorted
	sort.Sort(sort.Reverse(sort.StringSlice(s2)))
	fmt.Println(s2)
}
