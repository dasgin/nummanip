package calc

import (
	"errors"
	"github.com/fatih/color"
)

// Add and return sum of integers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		errorMessage := color.RedString("provide more than 2 numbers")
		return sum, errors.New(errorMessage)
	} 

	for _, num := range numbers {
		sum = sum + num
	}

	return sum, nil
}
