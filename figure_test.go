package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestReadMultipleFiles(t *testing.T) {
	// set the random seed to 0
	rand.Seed(0)

	// create the XOR representation patter to train the network
	/*
		X's size is 2400*1152
		Y's size is 2400*24
	*/

	X, Y := ReadMultipleFiles("ass2_processed_data")
	fmt.Println(Y[0])
	fmt.Println(len(Y))
	fmt.Println(len(X))
	fmt.Println(X[0])
}

func TestDataPartition(t *testing.T) {
	X, Y := ReadMultipleFiles("ass2_processed_data")
	fmt.Println(len(Y[0]))
	X_train, X_test, _, _ := DataPartition(X, Y, 0.80)
	fmt.Println(X_train[0])
	//fmt.Println(X_train[1])
	//fmt.Println(X_train[2])
	fmt.Println(len(X_test))
}

func TestImageTrain(t *testing.T) {
	X, Y := ReadMultipleFiles("ass2_processed_data")
	X_train, _, y_train, _ := DataPartition(X, Y, 0.80)
	net := CreateNetwork(1152, 200, 24, 0.1)
	ImageTrain(&net, X_train, y_train)
	save(net)
}

func TestImagePredict(t *testing.T) {
	X, Y := ReadMultipleFiles("ass2_processed_data")
	net := CreateNetwork(1152, 200, 24, 0.1)
	load(&net)
	ImagePredict(&net, X, Y)
	//fmt.Println(Y[101])

}
