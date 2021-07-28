package main

import "fmt"

func AmountOfElements(numbers []int) int {
	result := 0
	for _, value := range numbers {
		result += value
	}
	return result
}

func MiddleValue(numbers []float64) float64 {
	var amount float64
	var result float64
	for _, value := range numbers {
		amount += float64(value)
	}
	lenOfSlice := len(numbers)
	if lenOfSlice == 0 {
		lenOfSlice = 1
	}
	result = amount / float64(lenOfSlice)
	return result
}

func main() {
	fmt.Println(AmountOfElements([]int{10, -2, 43, 23, -11}))
	fmt.Println(MiddleValue([]float64{10, 2, 43, 23, 11}))
}
