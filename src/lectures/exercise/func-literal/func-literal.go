//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	lettersNum, digitsNum, spacesNum, punctuationMarksNum := 0, 0, 0, 0
	for _, line := range lines {
		for _, r := range line {
			if unicode.IsLetter(r) {
				lettersNum++
			} else if unicode.IsDigit(r) {
				digitsNum++
			} else if unicode.IsSpace(r) {
				spacesNum++
			} else if unicode.IsPunct(r) {
				punctuationMarksNum++
			}
		}
	}
	fmt.Printf("Letters: %d\n", lettersNum)
	fmt.Printf("Digits: %d\n", digitsNum)
	fmt.Printf("Spaces: %d\n", spacesNum)
	fmt.Printf("Punctuation marks: %d\n", punctuationMarksNum)
}
