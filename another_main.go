package main

import (
	"fmt"
)

func main() {
	var a int //this is a variable
	fmt.Println("Enter your Number: ")
	fmt.Scanln(&a)
	fmt.Println("the number you entered is ", a)
}
