package main

import (
	"math/rand"
	"time"
)

/*
This would separate the X and Y value into training (80%) and testing data (20%)
*/

func DataPartition(X [][]float64, Y []string) ([][]float64, [][]float64, []string, []string) {

	return nil, nil, nil, nil
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
