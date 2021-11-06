package main

import "gonum.org/v1/gonum/mat"

// Network contains the number of neurons in input layer, hidden layer and output layer
type Network struct {
	inputs        int
	hiddens       int
	outputs       int
	hiddenWeights *mat.Dense
	outputWeights *mat.Dense
	learningRate  float64
}

func CreateNetwork(input, hidden, output int, lr float64) Network {
	net := Network{
		inputs:       input,
		hiddens:      hidden,
		outputs:      output,
		learningRate: lr,
	}

	// random initialize the weights for input and hidden layers
	// S_j = sum(W_ji*a_i)
	// j is the number of neuron in hidden layer
	// i is the number of neuron in input layer
	//net.hiddenWeights = mat.NewDense(hidden, input, generateRandArray(hidden*input, float64(input)))
	//net.outputWeights = mat.NewDense(output, hidden, generateRandArray(output*hidden, float64(hidden)))

	return net
}
