package main

import "fmt"

func main() {
	createIterator := func() func() int {
		counter := 0
		return func() int {
			counter++
			return counter
		}
	}
	iterator := createIterator()
	fmt.Println(iterator())
	fmt.Println(iterator())
	fmt.Println(iterator())
	fmt.Println(iterator())
	fmt.Println(iterator())
	newIterator := createIterator()
	fmt.Println(newIterator())
	fmt.Println(newIterator())
	fmt.Println(newIterator())
	fmt.Println(newIterator())
	fmt.Println(newIterator())
	fmt.Println(iterator())
	fmt.Println(iterator())
	fmt.Println(iterator())
	fmt.Println(newIterator())
	fmt.Println(newIterator())
	fmt.Println(newIterator())
}
