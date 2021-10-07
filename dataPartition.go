package main

import (
	"math/rand"
	"time"
)

/*
This would separate the X and Y value into training (80%) and testing data (20%)
*/

func DataPartition(X [][]float64, Y []float64, ratio float64) ([][2][]float64, [][2][]float64, [][]float64, [][]float64, [][]float64, [][]float64) {

	allDataNum := len(X)
	trainNum := int(float64(allDataNum) * ratio)
	testNum := allDataNum - trainNum

	allIndex := make([]int, 0)
	for i := 0; i < len(X); i++ {
		allIndex = append(allIndex, i)
	}

	// shuffle all the index
	Shuffle(allIndex)

	// divide into training and testing set based on the previous division
	trainIndex := allIndex[:trainNum]
	testIndex := allIndex[trainNum:]

	X_train := make([][]float64, 0)
	X_test := make([][]float64, 0)
	y_train := make([][]float64, 0)
	y_test := make([][]float64, 0)

	inputTrain := make([][2][]float64, trainNum)
	for idx, val := range trainIndex {
		X_train = append(X_train, X[val])
		y_train = append(y_train, []float64{Y[val]})
		inputTrain[idx][0] = X[val]
		inputTrain[idx][1] = []float64{Y[val]}
	}
	inputTest := make([][2][]float64, testNum)
	for idx, val := range testIndex {
		X_test = append(X_test, X[val])
		y_test = append(y_test, []float64{Y[val]})
		inputTrain[idx][0] = X[val]
		inputTrain[idx][1] = []float64{Y[val]}
	}

	//fmt.Println(inputTrain[0])
	//fmt.Println(inputTest[0])
	//inputTrain := make([][]float64,1)
	//inputTrain[0] =X[0]
	//inputTrain[1] =y_train[0]
	//fmt.Println(inputTrain)
	//fmt.Println(X[0])
	//fmt.Println(y_train[0])
	//fmt.Println(y_train)
	//fmt.Println("==========================")
	//fmt.Println(y_test)
	//fmt.Println(X_train)
	//fmt.Println(train_Index)
	//fmt.Println(len(train_Index))
	//fmt.Println(len(test_Index))
	//fmt.Println(test_Index)

	return inputTrain, inputTest, X_train, X_test, y_train, y_test
}

func Shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
