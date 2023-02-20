//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func printInfo(prods []Product) {
	fmt.Println("Last item:", prods[len(prods)-1])
	fmt.Println("Length:", len(prods))
	sum := 0.0
	for i := 0; i < len(prods); i++ {
		item := prods[i]
		sum += item.price
	}
	fmt.Println("Sum:", sum)
}

func main() {
	shoppintList := []Product{{"carrots", 1.50}, {"milk", 2}, {"oranges", 3}}
	printInfo(shoppintList)
	shoppingList := append(shoppintList, Product{"apples", 2})
	printInfo(shoppingList)
}
