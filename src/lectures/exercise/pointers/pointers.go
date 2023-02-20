//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type SecuredProduct struct {
	name  string
	state bool
}

func activateTag(securedProduct *SecuredProduct) {
	securedProduct.state = true
}

func deactivateTag(securedProduct *SecuredProduct) {
	securedProduct.state = false
}

func checkout(securedProducts *[]SecuredProduct) {
	for _, securedProduct := range *securedProducts {
		securedProduct.state = false
	}
}

func main() {
	var products []SecuredProduct
	fmt.Println(products)
	products = append(products, SecuredProduct{name: "Apple", state: true},
		SecuredProduct{name: "Orange", state: true},
		SecuredProduct{name: "Banana", state: true},
		SecuredProduct{name: "Mango", state: false})
	activateTag(&products[len(products)-1])
	fmt.Println(products)
	deactivateTag(&products[0])
	fmt.Println(products)
	checkout(&products)
	fmt.Println(products)
}
