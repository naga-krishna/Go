package main

import (
	"fmt"

	child "program2/child"
	father "program2/parent"
)

func main() {
	f := father.Father{}
	s := child.Son{}
	fmt.Println(f.Data("Nag"))
	fmt.Println(s.Data("Krishna"))
}
