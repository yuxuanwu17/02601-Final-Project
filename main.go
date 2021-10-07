package main

import (
	"fmt"
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
	_, _, y_train, y_test := DataPartition(X, Y, 0.8)
	fmt.Println(y_train)
	fmt.Println(y_test)
	//
	//fmt.Println(len(x_train))
	//fmt.Println(len(y_train))

	//inputData := make([][][]int, 0)
	//for i := 0; i < len(x_train); i++ {
	//inputData[i] = append(x_train[i],y_train[i])
	//}

	// patterns

	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}

	fmt.Println(patterns[0])
	fmt.Println(patterns[0][0])
	// instantiate the Feed Forward
	ff := &gobrain.FeedForward{}

	// initialize the Neural Network;
	// the networks structure will contain:
	// 2 inputs, 2 hidden nodes and 1 output.
	ff.Init(2, 2, 1)

	// train the network using the XOR patterns
	// the training will run for 1000 epochs
	// the learning rate is set to 0.6 and the momentum factor to 0.4
	// use true in the last parameter to receive reports about the learning error
	ff.Train(patterns, 1000, 0.6, 0.4, true)
}

//DataPartition(X, Y, 0.8)
