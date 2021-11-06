package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func Test_dot(t *testing.T) {
	// Initialize two matrices, a and b.
	a := mat.NewDense(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})
	b := mat.NewDense(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})
	fmt.Println(a)
	fmt.Println(b)
	dot(a, b)
}
