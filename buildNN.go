package main

import (
	"gonum.org/v1/gonum/mat"
)

// Network contains the number of neurons in input layer, hidden layer and output layer
type Network struct {
	inputs        int
	hiddens       int
	outputs       int
	hiddenWeights *mat.Dense
	outputWeights *mat.Dense
	lr            float64
	squaredErr    float64
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

// Train the neuron network given the input and target data
func (net *Network) Train(input []float64, target []float64) {
	// feedforward network
	finalOutputMat, hiddenOutputMat, inputMat := net.FeedForward(input)

	// error calculation
	outputError, hiddenError := net.ErrorCalculation(finalOutputMat, target)

	// backpropagation
	net.BackPropagation(outputError, hiddenError, finalOutputMat, hiddenOutputMat, inputMat)
}

// FeedForward would conduct the feedforward, return the finalOutputMat, ready to compare Error with actual target
func (net *Network) FeedForward(input []float64) (mat.Matrix, mat.Matrix, mat.Matrix) {
	inputMat := mat.NewDense(len(input), 1, input)
	hiddenInputMat := dot(net.hiddenWeights, inputMat)
	hiddenOutputMat := applySigmoid(sigmoid, hiddenInputMat)
	finalInputMat := dot(net.outputWeights, hiddenOutputMat)
	finalOutputMat := applySigmoid(sigmoid, finalInputMat)
	return finalOutputMat, hiddenOutputMat, inputMat
}

// ErrorCalculation would return the error in each layer
func (net *Network) ErrorCalculation(finalOutputMat mat.Matrix, target []float64) (mat.Matrix, mat.Matrix) {
	targetMat := mat.NewDense(len(target), 1, target)
	outputError := subtract(targetMat, finalOutputMat)
	squaredErr := obtainMSE(finalOutputMat, targetMat)
	net.squaredErr += squaredErr
	hiddenError := dot(net.outputWeights.T(), outputError)
	return outputError, hiddenError
}

// BackPropagation would conduct the backpropagation during the network
func (net *Network) BackPropagation(outputError, hiddenError, finalOutputMat, hiddenOutputMat, inputMat mat.Matrix) {
	// part 1: the target matrix - output matrix
	// outputError
	// part 2: the derivative of the final output with respect to the summation of k
	d_ok_dSumk := sigmoidDerivative(finalOutputMat)
	d_ok_dSumk_hidden := sigmoidDerivative(hiddenOutputMat)

	// part 1 element multiply of part 2
	part12 := multiply(outputError, d_ok_dSumk)
	part12_hidden := multiply(hiddenError, d_ok_dSumk_hidden)
	// part 3: the gradient of summation with respect to the output weight wjk, which is also the input
	// dot product of the transpose of the previous output, necessary for multiplying layers
	deltaWeights := dot(part12, hiddenOutputMat.T())
	deltaWeights_hidden := dot(part12_hidden, inputMat.T())

	net.outputWeights = add(net.outputWeights, scale(net.lr, deltaWeights)).(*mat.Dense)

	net.hiddenWeights = add(net.hiddenWeights, scale(net.lr, deltaWeights_hidden)).(*mat.Dense)
}

// Predict the input given input data, 1D array
func (net Network) Predict(input []float64) mat.Matrix {
	inputMat := mat.NewDense(len(input), 1, input) // convert the input to vector format
	hiddenInputMat := dot(net.hiddenWeights, inputMat)
	hiddenOutputMat := applySigmoid(sigmoid, hiddenInputMat)
	finalInputMat := dot(net.outputWeights, hiddenOutputMat)
	finalOutputMat := applySigmoid(sigmoid, finalInputMat)
	return finalOutputMat
}
