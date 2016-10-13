package main

import (
	"testing"
)

func TestTrecimate(t *testing.T) {
	var expected = 1
	var actual = trecimate(3)
	if (actual != expected) {
		t.Errorf("canDivideByThree: Expected %d", expected)
		t.Errorf("canDivideByThree: Actual %d", actual)
	}

	expected = 3
	actual = trecimate(10)

	if (actual != expected) {
		t.Errorf("subtractOneAndDivideByThree: Expected %d", expected)
		t.Errorf("subtractOneAndDivideByThree: Actual %d", actual)
	}

	expected = 12
	actual = trecimate(35)

	if (actual != expected) {
		t.Errorf("addOneAndDivideByThree: Expected %d", expected)
		t.Errorf("addOneAndDivideByThree: Actual %d", actual)
	}

}

