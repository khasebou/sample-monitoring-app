package main

import (
	"fmt"
)

func DotProduct(x, y []int) (int, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("vector size mismatch, x is %d elements, y is %d elements",
			len(x), len(y))
	}
	sum := 0
	for i := range x {
		sum += x[i] * y[i]
	}
	return sum, nil
}
