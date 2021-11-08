package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func Test_dot(t *testing.T) {
	//Initialize two matrices, a and b.
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
	fmt.Println(dot(a, b))

	//a := mat.NewDense(3, 3, []float64{
	//	1.5, 0.4, -0.3,
	//	0.4, 1.4, 0.8,
	//	-0.3, 0.8, 2.3,
	//})
	//b := mat.NewDense(3, 1, []float64{
	//	5.9,
	//	6.9,
	//	3.7,
	//})
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(dot(a, b))

}
func TestSubtract(t *testing.T) {
	a := mat.NewDense(2, 2, []float64{
		1, 2,
		4, 5,
	})
	b := mat.NewDense(2, 2, []float64{
		7, 8,
		9, 10,
	})
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(subtract(a, b))
}
func TestSigmoid(t *testing.T) {
	fmt.Println(applySigmoid(sigmoid, mat.NewDense(2, 1, []float64{3, 3})))
	fmt.Println(sigmoid(2, 1, 3))
}
func TestNetwork_FeedForward(t *testing.T) {
	net := CreateNetwork(2, 2, 1, 0.1)
	net.hiddenWeights = mat.NewDense(2, 2, []float64{0.0, 1.0, 0.0, 1.0})
	net.outputWeights = mat.NewDense(1, 2, []float64{0.0, 1.0})
	output, _, _ := net.FeedForward([]float64{2, 3})
	fc := mat.Formatted(output, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("c = %v", fc)
}
func TestInitialWeights(t *testing.T) {
	fmt.Println(initialWeights(10))
}
