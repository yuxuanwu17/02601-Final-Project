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

	X_train, X_test, _, _ := DataPartition(X, Y, 0.80)
	fmt.Println(X_train[0])
	//fmt.Println(X_train[1])
	//fmt.Println(X_train[2])
	fmt.Println(len(X_test))
}