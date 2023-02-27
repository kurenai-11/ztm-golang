//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func validateString(s string) int {
	if s == "" {
		return 0
	}
	number, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return number
}

func processFile(filename string, sums *[]int, index int, wg *sync.WaitGroup) {
	sumsSlice := *sums
	numbersSum := 0
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentNumber := scanner.Text()
		validatedNumber := validateString(currentNumber)
		numbersSum += validatedNumber
	}
	sumsSlice[index] = numbersSum
	wg.Done()
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	totalSum := 0
	sums := make([]int, len(files))
	wg := &sync.WaitGroup{}
	for index, file := range files {
		wg.Add(1)
		go processFile(file, &sums, index, wg)
	}
	wg.Wait()
	for _, sum := range sums {
		totalSum += sum
	}
	fmt.Println("Sum is:", totalSum)
}
