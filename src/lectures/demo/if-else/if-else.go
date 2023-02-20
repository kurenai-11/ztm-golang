package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 2, 2, 3
	if quiz1 > quiz2 {
		fmt.Println("quiz1>quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz1<quiz2")
	} else {
		fmt.Println("quizzes 1 and 2 are equal")
	}

	if average(quiz1, quiz2, quiz3) > 1.99 {
		fmt.Println("ok")
		fmt.Println(average(quiz1, quiz2, quiz3))
	}
}
