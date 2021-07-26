package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Print("Enter string 1: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter string 2: ")
	fmt.Scanln(&str2)
	fmt.Println("Combined Strings: ", str1+str2)
}
