package main

import (
	"fmt"

	"github.com/mvonbodun/go-test/stringutil"
)

// test structure
type VB struct {
	x    int
	name string
}

func main() {

	fmt.Println("Before initializing struct")
	nameStruct := VB{1, "Michael"}
	x := 1
	var y int
	y = x + 1
	fmt.Println(nameStruct)
	fmt.Println(y)
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("This is vb2")
	fmt.Println("Here is the 3rd line")
	fmt.Println("Here is the 4th line")
	fmt.Println("The 5th line")
	fmt.Println("The 6th line")
	fmt.Println("The 8th line")
	fmt.Println("The 9th line")
	// The 7th line
	// Next up
}
