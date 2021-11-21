package main

import (
	"fmt"
)

func ImageTrain(net *Network, X, Y [][]float64) {
	for epoch := 0; epoch < 5; epoch++ {
		// iterate first 10 picture
		for i := 0; i < len(Y); i++ {
			net.Train(X[i], Y[i])
		}
	}
}

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
		}
	}
	fmt.Println("The number of test case:", len(X_test))
	fmt.Println("score is:", score)
	total := len(X_test)
	accuracy := float64(score) / float64(total)
	fmt.Println("accuracy is:", accuracy)
}
