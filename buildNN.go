package main

import "gonum.org/v1/gonum/mat"

// Network contains the number of neurons in input layer, hidden layer and output layer
type Network struct {
	inputs        int
	hiddens       int
	outputs       int
	hiddenWeights *mat.Dense
	outputWeights *mat.Dense
	lr            float64
}

// CreateNetwork accept the dimension of each input figure 1152
func CreateNetwork(input, hidden, output int, lr float64) Network {
	net := Network{
		inputs:  input,
		hiddens: hidden,
		outputs: output,
		lr:      lr,
	}

	// S_j = sum(W_ji*a_i)  j : hidden layer, i : input layer
	// random initialize the weights for input and hidden layers
	net.hiddenWeights = mat.NewDense(hidden, input, initialWeights(hidden*input))
	net.outputWeights = mat.NewDense(output, hidden, initialWeights(output*hidden))

	return net
}

func (net *Network) Train(input []float64, target []float64) {
	// feedforward network
	finalOutputMat := net.FeedForward(input)

	// error calculation
	outputError, hiddenError := net.ErrorCalculation(finalOutputMat, target)

	// backpropagation

}

// FeedForward would conduct the feedforward, return the finalOutputMat, ready to compare Error with actual target
func (net *Network) FeedForward(input []float64) mat.Matrix {
	inputMat := mat.NewDense(len(input), 1, input)
	hiddenInputMat := dot(net.hiddenWeights, inputMat)
	hiddenOutputMat := applySigmoid(sigmoid, hiddenInputMat)
	finalInputMat := dot(net.outputWeights, hiddenOutputMat)
	finalOutputMat := applySigmoid(sigmoid, finalInputMat)
	return finalOutputMat
}

// ErrorCalculation would return the error in each layer
func (net *Network) ErrorCalculation(finalOutputMat mat.Matrix, target []float64) (mat.Matrix, mat.Matrix) {
	targetMat := mat.NewDense(len(target), 1, target)
	outputError := subtract(targetMat, finalOutputMat)
	hiddenError := dot(net.outputWeights.T(), outputError)
	return outputError, hiddenError
}
