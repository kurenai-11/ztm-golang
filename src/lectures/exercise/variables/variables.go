//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favoriteColor = "red"
	fmt.Println("my favorite color is", favoriteColor)
	const birthYear, yearNow = 2002, 2023
	fmt.Println("my birth year is", birthYear)
	age := yearNow - birthYear
	fmt.Println("my age is", age)
	var (
		first = 'J'
		last  = 'K'
	)
	fmt.Println("my initials are", first, last)
	ageInDays := 365 * age
	fmt.Println("my age in days is", ageInDays)
}
