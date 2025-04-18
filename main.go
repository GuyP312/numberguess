package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func randInt(cap int) int {
	seed := time.Now().UnixNano()      // get current timestamp in nanosecond
	source := rand.NewSource(seed)     // create new source of randomess from seed value which is timestamp
	randInstance := rand.New(source)   // create new random number generator
	goal := randInstance.Intn(cap + 1) // generate random number between 0 and cap + 1
	return goal
}

func main() {
	cap := 0
	fmt.Println("Hi, this is a number guessing game. Please enter the upper bound:")
	fmt.Scan(&cap)

	goal := randInt(cap)
	guess := 0

	maximumAttempts := int(math.Ceil(math.Log2(float64(cap)))) // log2 the cap and round up to the nearest integer
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
