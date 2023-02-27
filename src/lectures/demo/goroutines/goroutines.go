package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune
	capitalize := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
	}
	for _, r := range data {
		go capitalize(r)
	}
	time.Sleep(200)
	for i := 0; i < len(data); i++ {
		char := data[i]
		fmt.Printf("%c", char)
	}
	fmt.Printf("\n")
}
