package main

import (
	"flag"
	"fmt"
)

func main() {
	// X size is 2400*1152, Y size is 2400*24
	X, Y := ReadMultipleFiles("ass2_processed_data")
	X_train, X_test, y_train, y_test := DataPartition(X, Y, 0.80)
	net := CreateNetwork(1152, 200, 24, 0.1)
	option := flag.String("option", "", "Select train/predict to train or predict the neural network")
	file := flag.String("file", "ass2_processed_data/C1.jpeg", "File name of any PNG file in the ass2_processed_data")

	flag.Parse()
	switch *option {
	case "train":
		fmt.Println("You are ready to train the model")
		ImageTrain(&net, X_train, y_train)
		save(net)
	case "predict":
		fmt.Println("You are going to predict the model based on the input file")
		load(&net)
		ImagePredict(&net, X_test, y_test)

	case "":
		//fmt.Println("Please select one option, you can use -help for more details")
	}

	if *file != "" {
		//http.HandleFunc("/predict", PredictSingleFigure)
		//err := http.ListenAndServe(":9090", nil)
		//if err != nil {
		//	fmt.Printf("http serve failed, err:%v\n", err)
		//}
		fig := ReadSingleFile(*file)
		load(&net)
		fmt.Println("The predicted label is:", TokenToLabel(SingleImagePredict(&net, fig)))
		fmt.Println("The input image label is:", ObtainLabelFromString(*file))
	}

}

//func PredictSingleFigure(w http.ResponseWriter, r *http.Request) {
//	//b ,_ :=ioutil.ReadFile("web/web.txt")
//	_, _ = fmt.Fprintf(w, "<h1>Hello world</h1>\n" +
//		"<button id= 'b1'>点我</button>\n"+
//		"<img src='plot/python_cm.png'>")
//}
