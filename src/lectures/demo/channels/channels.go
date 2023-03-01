package main

import "fmt"

func main() {
	channel := make(chan int, 2)
	channel <- 1
	channel <- 2
	go func() {
		channel <- 3
	}()
	first := <-channel
	second := <-channel
	third := <-channel
	fmt.Println(first, second, third)
}
