//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func isNewborn(age int) bool {
	return age == 0
}

func isToddler(age int) bool {
	return age >= 1 && age <= 3
}

func isChild(age int) bool {
	return age >= 4 && age <= 12
}

func isTeenager(age int) bool {
	return age >= 13 && age <= 17
}

func isAdult(age int) bool {
	return age >= 18
}

func main() {
	switch age := 122; {
	case isNewborn(age):
		fmt.Println("newborn")
	case isToddler(age):
		fmt.Println("toddler")
	case isChild(age):
		fmt.Println("child")
	case isTeenager(age):
		fmt.Println("teenager")
	case isAdult(age):
		fmt.Println("adult")
	}
}
