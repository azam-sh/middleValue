package main

import "testing"

func TestAmountOfElements(t *testing.T) {
	tests := []struct {
		Name       string
		Numbers    []int
		AmountWant int
	}{
		{"All numbers are positive", []int{10, 2, 43, 23, 11}, 89},
		{"All numbers are negative", []int{-10, -2, -43, -23, -11}, -89},
		{"Positive and negative numbers", []int{10, -2, -43, 23, -11}, -23},
	}
	for _, test := range tests {
		if amount := AmountOfElements(test.Numbers); amount != test.AmountWant {
			t.Error("Test Failed: {} inputted, expected, received: {}", test.Numbers, test.AmountWant, amount)
		}
	}
}

func TestMiddleValue(t *testing.T) {
	tests := []struct {
		Name            string
		Numbers         []float64
		MiddleValueWant float64
	}{
		{"All numbers are positive", []float64{10, 2, 43, 23, 11}, 17.8},
		{"All numbers are negative", []float64{-10, -2, -43, -23, -11}, -17.8},
		{"Positive and negative numbers", []float64{10, -2, -43, 23, -11}, -4.6},
		{"Positive and negative numbers", []float64{}, 0},
	}
	for _, test := range tests {
		if result := MiddleValue(test.Numbers); result != float64(test.MiddleValueWant) {
			t.Error("Test Failed: {} inputted, expected, received: {}", test.Numbers, test.MiddleValueWant, result)
		}
	}
}
