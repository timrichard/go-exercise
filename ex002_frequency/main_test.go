package main

import (
	"testing"
	"fmt"
)

//func assert(actual int, expected int, errorMessage string, errorReporter *testing.T) {
//	if (actual != expected) {
//		errorReporter.Errorf(errorMessage, expected, actual)
//	}
//}

func TestAnalyseFrequency(t *testing.T) {
	var frequency map[int]int
	var uniques []int

	nums := []int{10, 5, 2, 10}

	frequency, uniques = AnalyseFrequency(nums)
	fmt.Println(frequency)
	fmt.Println(uniques)
}
