package calculator_test

import (
	"hello/calculator"
	"testing"
)


var dataTest1 = [][]int{
	{1,2,3},
	{4,5,6},
	{9,8,9},
}

var dataTest2 = [][]int{
	{-1, 2, 3},
	{4, -5, 6},
	{9, 8, -9},
}

var dataTest3 = [][]int{
	{1, 1, 1, 1, 1, 1, 1},
	{3, 3, 3, 3, 3, 3, 3},
	{0, 0, 4, 5, 4, 0, 0},
	{0, 0, 0, 5, 0, 0, 0},
	{0, 0, -4, 5, -4, 0, 0},
	{0, -3, 0, 5, 0, -3, 0},
	{-1, 0, 0, 5, 0, 0, -1},
}

func TestData1Should32(t *testing.T) {
	expected := 2

	result := calculator.Calculate(dataTest1)
	if result != expected {
		t.Errorf("expected %q but result %q\n", expected, result)
	}
}

func TestData2Should22(t *testing.T) {
	expected := 22

	result := calculator.Calculate(dataTest2)
	if result != expected {
		t.Errorf("expected %q but result %q\n", expected, result)
	}
}

func TestData3Should0(t *testing.T) {
	expected := 0

	result := calculator.Calculate(dataTest3)
	if result != expected {
		t.Errorf("expected %q but result %q\n", expected, result)
	}
}