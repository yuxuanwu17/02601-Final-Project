package main

import "fmt"

func main() {
	//ReadSingleFile("ass2_processed_data/A1.jpeg")
	X, Y := ReadMultipleFiles("ass2_processed_data")
	fmt.Println()
	fmt.Println(len(X))
	fmt.Println(len(X[0]))
	fmt.Println(len(Y))
	//x_train, x_test, y_train, y_test := DataPartition(X, Y)
	//fmt.Println(len(x_train))
	//fmt.Println(len(x_test))
	//fmt.Println(len(y_train))
	//fmt.Println(len(y_test))
}
