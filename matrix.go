package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

// The dot product between two matrix, used in the weight calculation between layers
func dot(a, b mat.Matrix) mat.Matrix {
	// Take the matrix product of a and b and place the result in c.
	var c mat.Dense
	c.Mul(a, b)
	return &c
}

// subtract function is to subtract the element between two matrix
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

// sigmoid activation function, the last layer activation layer
func sigmoid(r, c int, x float64) float64 {
	return 1.0 / (1 + math.Exp(-1*x))
}

// multiply is to conduct the element wise multiplication between two matrix m and n
func multiply(m, n mat.Matrix) mat.Matrix {
	r, c := m.Dims()
	o := mat.NewDense(r, c, nil)
	o.MulElem(m, n)
	return o
}

// scale would scale the input matrix m based on the float s
func scale(s float64, m mat.Matrix) mat.Matrix {
	r, c := m.Dims()
	o := mat.NewDense(r, c, nil)
	o.Scale(s, m)
	return o
}

// add would add m and n element-wise
func add(m, n mat.Matrix) mat.Matrix {
	r, c := m.Dims()
	o := mat.NewDense(r, c, nil)
	o.Add(m, n)
	return o
}

// calculate the derivative of sigmoid function, used in the backpropagation
func sigmoidDerivative(m mat.Matrix) mat.Matrix {
	rows, _ := m.Dims()
	o := make([]float64, rows)
	// create a one matrix
	for i := range o {
		o[i] = 1
	}
	ones := mat.NewDense(rows, 1, o)
	return multiply(m, subtract(ones, m)) // m * (1 - m)
}

// matrixPrint would print a Gonum matrix in a pretty way
func matrixPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

// Apply calculates the MSE given predicted and actual labels
func obtainMSE(prediction, actual mat.Matrix) float64 {
	rows, _ := prediction.Dims()
	err := 0.0
	for i := 0; i < rows; i++ {
		diff := prediction.At(i, 0) - actual.At(i, 0)
		err += math.Pow(diff, 2)
	}
	return err / float64(rows)
}
