package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

func main() {
	// X size is 2400*1152, Y size is 2400*24
	X, Y := ReadMultipleFiles("data")
	X_train, X_test, y_train, y_test := DataPartition(X, Y, 0.80)
	net := CreateNetwork(1152, 200, 24, 0.1)
	option := flag.String("option", "predict", "Select train/predict/singleImagePred to train or predict the neural network")
	file := flag.String("file", "", "File name of any PNG file in the data, e.g. data/C1.jpeg")
	epoch := flag.String("epoch", "5", "Define the number of epoch used in the training")

	flag.Parse()
	switch *option {
	case "train":
		fmt.Println("You are ready to train the model")
		numEpoch, err := strconv.Atoi(*epoch)
		if err != nil {
			log.Fatal("Invalid input number for epoch number, the error is: ", err)
		}
		fmt.Println("The number of epochs used in the training: ", numEpoch)
		ImageTrain(&net, X_train, y_train, numEpoch)
		save(net)
	case "predict":
		fmt.Println("You are going to predict the model based on the input file")
		dir, _ := ioutil.ReadDir("plot")
		for _, d := range dir {
			os.RemoveAll(path.Join([]string{"plot", d.Name()}...))
		}
		load(&net)
		ImagePredict(&net, X_test, y_test)

	case "singleImagePred":
		fig := ReadSingleFile(*file)
		load(&net)
		fmt.Println("The predicted label is:", TokenToLabel(SingleImagePredict(&net, fig)))
		fmt.Println("The input image label is:", ObtainLabelFromString(*file))
	}

}
