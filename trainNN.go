package main

import "fmt"

func ImageTrain(net *Network, X, Y [][]float64) {
	for epoch := 0; epoch < 5; epoch++ {
		// iterate first 10 picture
		for i := 0; i < len(Y); i++ {
			net.Train(X[i], Y[i])
		}
	}
}

func ImagePredict(net *Network, X_test, y_test [][]float64) {
	outputs := net.Predict(X_test[101])
	//fmt.Println(outputs)
	best := 0
	highest := 0.0
	score := 0
	for i := 0; i < net.outputs; i++ {
		if outputs.At(i, 0) > highest {
			best = i
			highest = outputs.At(i, 0)
		}
	}
	if y_test[101][best] == 1 {
		score++
	}
	fmt.Println(best)
	fmt.Println(y_test[101])
	fmt.Println("score is ", score)
}
