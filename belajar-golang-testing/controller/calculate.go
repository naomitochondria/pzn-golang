package controller

import "errors"

func CalculateSum(numbers ...int) (sum int) {
	for _, n := range numbers {
		sum += n
	}

	return
}

func CalculateDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero!")
	}

	return a / b, nil
}
