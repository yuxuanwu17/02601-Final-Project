package main

import (
	"fmt"
	"sort"
)

func main() {
	/*X's size is 2400*1152 */
	//X, Y := ReadMultipleFiles("ass2_processed_data")
	//x_train, x_test, y_train, y_test := DataPartition(X, Y)
	//DataPartition(X, Y)
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 10; i++ {
		Shuffle(test)
		fmt.Println(test)
		sort.Ints(test)
		fmt.Println(test)
		fmt.Println("====================")
	}
	//fmt.Println(len(X))
	//fmt.Println(len(X[0]))
	//fmt.Println(len(Y))
	//fmt.Println(len(x_train))
	//fmt.Println(len(x_test))
	//fmt.Println(len(y_train))
	//fmt.Println(len(y_test))
}
