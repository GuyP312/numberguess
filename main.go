package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func randInt(cap int) int {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	randInstance := rand.New(source)
	goal := randInstance.Intn(cap + 1)
	return goal
}

func main() {
	cap := 0
	fmt.Println("Hi, this is a number guessing game. Please enter the upper bound:")
	fmt.Scan(&cap)

	goal := randInt(cap)
	guess := 0

	maximumAttempts := int(math.Ceil(math.Log2(float64(cap))))
	for try := 0; try < maximumAttempts; try++ {
		fmt.Print("Please Enter your guess:")
		fmt.Scan(&guess)
		switch {
		case guess == goal:
			fmt.Println("You got it correct!")
			return
		case guess < goal:
			fmt.Println("Too small!")
		default:
			fmt.Println("Too large")
		}
	}
	fmt.Println("You reach your maximum limit! GG")
}
