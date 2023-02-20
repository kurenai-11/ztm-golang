//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Rectangle struct {
	side int
}

func getArea(rect Rectangle) int {
	return int(math.Pow(float64(rect.side), 2))
}

func getPerimeter(rect Rectangle) int {
	return rect.side * 4
}

func main() {
	fmt.Println(rand.Intn(5))
	r1 := Rectangle{8}
	fmt.Println("r1's area:", getArea(r1))
	fmt.Println("r1's perimeter:", getPerimeter(r1))
	r1 = Rectangle{5}
	fmt.Println("r1's area:", getArea(r1))
	fmt.Println("r1's perimeter:", getPerimeter(r1))
}
