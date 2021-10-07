package main

func main() {
	/*X's size is 2400*1152 */
	X, Y := ReadMultipleFiles("ass2_processed_data")
	//x_train, x_test, y_train, y_test := DataPartition(X, Y,0.8)
	DataPartition(X, Y, 0.8)

}
