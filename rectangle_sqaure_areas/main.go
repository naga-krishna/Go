package main

import (
	"fmt"
)

func main() {
	var reclen, recwid, squlen int
	fmt.Println("---------Rectangle---------")
	fmt.Print("Enter Length of the Rectangle: ")
	fmt.Scanln(&reclen)
	fmt.Print("Enter width of the Rectangle: ")
	fmt.Scanln(&recwid)
	fmt.Println("Area of Rectangle: ", reclen*recwid)
	fmt.Println("---------Square------------")
	fmt.Print("Enter Length of the Square: ")
	fmt.Scanln(&squlen)
	fmt.Println("Area of Square:", squlen*squlen)
}
