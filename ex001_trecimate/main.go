package main

import (
	"fmt"
	"math/rand"
	"time"
)

var attempt int = 1;

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var runningTotal int = rand.Int()

	for notYetOne := true; notYetOne; notYetOne = (runningTotal != 1) {
		runningTotal = trecimate(runningTotal)
	}
}

func trecimate(currentValue int) int {
	fmt.Printf("Attempt %d. ", attempt)
	attempt = attempt + 1
	fmt.Printf("Value was %d. ", currentValue)

	if (currentValue - 1) % 3 == 0 {
		currentValue = currentValue - 1
		fmt.Printf("Subtracting one. Value now %d. ", currentValue)
	} else if (currentValue + 1) % 3 == 0 {
		currentValue = currentValue + 1
		fmt.Printf("Adding one. Value now %d. ", currentValue)
	}

	currentValue = currentValue / 3
	fmt.Printf("Dividing by three. Value is now %d.\n", currentValue)
	return currentValue
}