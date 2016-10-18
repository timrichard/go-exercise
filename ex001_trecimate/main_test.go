package main

import (
	"testing"
)

func assert(actual int, expected int, errorMessage string, errorReporter *testing.T) {
	if (actual != expected) {
		errorReporter.Errorf(errorMessage, expected, actual)
	}
}

func TestTrecimate(t *testing.T) {

	assert(trecimate(3), 1, "Divide by three: Expected %d, got %d", t)
	assert(trecimate(10), 3, "Subtract one, divide by three: Expected %d, got %d", t)
	assert(trecimate(35), 12, "Add one, divide by three: Expected %d, got %d", t)

}

