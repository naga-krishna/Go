package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("String Functions: ")
	fmt.Println(strings.Compare("Naga", "Krishna"))
	fmt.Println(strings.Contains("Naga", "ag"))
	fmt.Println(strings.ContainsAny("Naga Krishna", "a"))
	fmt.Println(strings.HasPrefix("Naga Krishna", "K"))
	fmt.Println(strings.Join([]string{"Naga", "Krishna"}, " "))
	fmt.Println(strings.Repeat("Naga Krishna ", 2))
	fmt.Println(strings.Replace("Naga Krishna", "Naga", "Ganta", 1))
	fmt.Println(strings.Split("Naga,Krishna,Ganta", ","))
	fmt.Println(strings.ToLower("NAGA KRISHNA"))
	fmt.Println(strings.ToUpper("naga Krishna"))
	fmt.Println(strings.TrimLeft("!!!!!Naga Krishna", "!a"))
}
