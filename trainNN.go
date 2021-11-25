package main

import (
	"fmt"
)

// ImageTrain would train the input data and epoch would specify the round of training
func ImageTrain(net *Network, X_train, y_train [][]float64, epoch int) {
	for num := 0; num < epoch; num++ {
		for i := 0; i < len(y_train); i++ {
			net.Train(X_train[i], y_train[i])
		}
		fmt.Print("epoch ", num+1, " finished!  ")
		sumError := net.squaredErr
		numData := float64(len(y_train))
		fmt.Println("MSE:", sumError/numData)
		net.squaredErr = 0
	}
}

// ImagePredict would predict the input data and normally would be the X_test data
func ImagePredict(net *Network, X_test, y_test [][]float64) {
	score := 0
	for i := 0; i < len(X_test); i++ {
		outputs := net.Predict(X_test[i])
		best := 0
		highest := 0.0
		for num := 0; num < net.outputs; num++ {
			if outputs.At(num, 0) > highest {
				best = num
				highest = outputs.At(num, 0)
			}
		}
		if y_test[i][best] == 1 {
			score++
		} else {
			fmt.Println("===================")
			fmt.Println("The actual label is :", ObtainIndexFromArray(y_test[i]))
			fmt.Println("The predicted label index is :", TokenToLabel(best))
			fmt.Println("The actual image saved is: ")
			Mtx2Pixel(X_test[i], TokenToLabel(best))
		}
	}
	fmt.Println("======================Overall model analysis======================")
	fmt.Println("The number of test case:", len(X_test))
	fmt.Println("The number of perfect match is:", score)
	total := len(X_test)
	accuracy := float64(score) / float64(total)
	fmt.Println("accuracy is:", accuracy)
}

// SingleImagePredict would accept the input figure and predict the label
func SingleImagePredict(net *Network, fig []float64) int {
	outputs := net.Predict(fig)
	best := 0
	highest := 0.0
	matrixPrint(outputs)
	for num := 0; num < net.outputs; num++ {
		if outputs.At(num, 0) > highest {
			best = num
			highest = outputs.At(num, 0)
		}
	}
	return best
}
