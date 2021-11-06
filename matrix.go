package main

import (
	"gonum.org/v1/gonum/mat"
)

// The dot product between two matrix
// used in the weight calculation between layers
func dot(a, b mat.Matrix) mat.Matrix {
	// Take the matrix product of a and b and place the result in c.
	var c mat.Dense
	c.Mul(a, b)
	return &c
}

//func generateRandArray(size int,v float64) []float64 {
//
//}
