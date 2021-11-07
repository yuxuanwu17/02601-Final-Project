package main

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

// The dot product between two matrix
// used in the weight calculation between layers
func dot(a, b mat.Matrix) mat.Matrix {
	// Take the matrix product of a and b and place the result in c.
	var c mat.Dense
	c.Mul(a, b)
	return &c
}

func subtract(a, b mat.Matrix) mat.Matrix {
	var c mat.Dense
	c.Sub(a, b)
	return &c
}

// applySigmoid wrap the function in mat.NewDense Apply method, which would apply the sigmoid function to its matrix
func applySigmoid(sigmoid func(row, col int, x float64) float64, old mat.Matrix) mat.Matrix {
	r, c := old.Dims()
	newMat := mat.NewDense(r, c, nil)
	newMat.Apply(sigmoid, old)
	return newMat
}

// initialWeights would initial random weight between -0.5 to 0.5
func initialWeights(size int) []float64 {
	rand.Seed(33)
	max := 0.5
	min := -0.5
	weights := make([]float64, size)
	for i := 0; i < size; i++ {
		weights[i] = rand.Float64()*(max-min) + min
	}

	return weights
}

// sigmoid activation function
func sigmoid(r, c int, x float64) float64 {
	return 1.0 / (1 + math.Exp(-1*x))
}

func sigmoidDerivative() {

}
