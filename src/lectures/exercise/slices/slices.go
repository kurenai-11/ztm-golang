//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printInfo(parts []Part) {
	fmt.Println("start!")
	for i := 0; i < len(parts); i++ {
		element := parts[i]
		fmt.Println(element)
	}
	fmt.Println("end!")
}

func main() {
	assemblyLine := []Part{"1", "2", "3"}
	printInfo(assemblyLine)
	assemblyLine = append(assemblyLine, "4", "5")
	printInfo(assemblyLine)
	assemblyLine = assemblyLine[3:]
	printInfo(assemblyLine)
}
