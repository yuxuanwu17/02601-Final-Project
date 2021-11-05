package main

import (
	"fmt"
	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
)

/*
test the successful install of the golang tensorflow
*/
//func main() {
//	root := tg.NewRoot()
//	A := tg.NewTensor(root, tg.Const(root, [2][2]int32{{1, 2}, {-1, -2}}))
//	x := tg.NewTensor(root, tg.Const(root, [2][1]int64{{10}, {100}}))
//	b := tg.NewTensor(root, tg.Const(root, [2][1]int32{{-10}, {10}}))
//	Y := A.MatMul(x.Output).Add(b.Output)
//	// Please note that Y is just a pointer to A!
//
//	// If we want to create a different node in the graph, we have to clone Y
//	// or equivalently A
//	Z := A.Clone()
//	results := tg.Exec(root, []tf.Output{Y.Output, Z.Output}, nil, &tf.SessionOptions{})
//	fmt.Println("Y: ", results[0].Value(), "Z: ", results[1].Value())
//	fmt.Println("Y == A", Y == A) // ==> true
//	fmt.Println("Z == A", Z == A) // ==> false
//}

//func main() {
//	// set the random seed to 0
//	rand.Seed(0)
//
//	// create the XOR representation patter to train the network
//	/*
//		X's size is 2400*1152
//		Y's size is 2400*24
//	*/
//
//	_, Y := ReadMultipleFiles("ass2_processed_data")
//	fmt.Println(Y)
//	fmt.Println(len(Y))
//	fmt.Println(len(Y[0]))
//	//DataPartition(X, Y, 0.99)
//	//inputTrain, _, _, _, _, _ := DataPartition(X, Y, 0.99)
//	//_, _, _, x_test, _, y_test := DataPartition(X, Y, 0.99)
//
//}

/**/
func main() {
	model := tg.LoadModel("model/test", []string{"serve"}, nil)

	fakeInput, _ := tf.NewTensor([1][1][1920]float64{})
	results := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_inputs_input", 0): fakeInput,
	})
	fmt.Println(results)
	//fmt.Println(model)
	//OneHotEncoding("G")
}
