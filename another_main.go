package main

import (
	"fmt"
)

func main() {
	var firstnum int  //this is a variable
	var secondnum int //this is variable declaration in go...
	fmt.Println("Enter your first Number: ")
	fmt.Scanln(&firstnum)
	fmt.Println("Enter your second Number: ")
	fmt.Scanln(&secondnum)
	fmt.Println("the number you entered is ", firstnum)
	var sum = firstnum + secondnum
	fmt.Println("the sum of first number ", firstnum, " and second number ", secondnum, "is ", sum)
}
