package main

import "fmt"

func ImageTrain(net *Network, X, Y [][]float64) {
	for epoch := 0; epoch < 5; epoch++ {
		// iterate first 10 picture
		for i := 0; i < 100; i++ {
			net.Train(X[i], Y[i])
		}
	}
}

func ImagePredict(net *Network, X_test [][]float64) {
	outputs := net.Predict(X_test[101])
	fmt.Println(outputs)
	//best := 0
	//highest := 0.0
}
