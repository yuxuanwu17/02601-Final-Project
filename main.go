package main

import (
	"github.com/goml/gobrain"
	"math/rand"
)

func main() {
	// set the random seed to 0
	rand.Seed(0)

	// create the XOR representation patter to train the network
	/*X's size is 2400*1152 */
	X, Y := ReadMultipleFiles("ass2_processed_data")
	// ReadMultipleFiles("ass2_processed_data")
	inputTrain, _, _, _, _, _ := DataPartition(X, Y, 0.99)

	// patterns

	//patterns := [][][]float64{
	//	{{0, 0}, {0}},
	//	{{0, 1}, {1}},
	//	{{1, 0}, {1}},
	//	{{1, 1}, {0}},
	//}

	//fmt.Println(patterns[0])
	//fmt.Println(patterns[0][0])
	//fmt.Println(patterns[0][1])
	// instantiate the Feed Forward
	ff := &gobrain.FeedForward{}

	// initialize the Neural Network;
	// the networks structure will contain:
	// 2 inputs, 2 hidden nodes and 1 output.
	ff.Init(2, 100, 1)

	// train the network using the XOR patterns
	// the training will run for 1000 epochs
	// the learning rate is set to 0.6 and the momentum factor to 0.4
	// use true in the last parameter to receive reports about the learning error
	//ff.Train(patterns, 1000, 0.6, 0.4, true)
	ff.Train(inputTrain, 1000, 0.6, 0.4, true)
}

//DataPartition(X, Y, 0.8)
