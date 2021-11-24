package main

import (
	"fmt"
	"net/http"
)

func webHandler() {
	http.HandleFunc("/predict", PredictSingleFigure)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve failed, err:%v\n", err)
	}
}

func PredictSingleFigure(w http.ResponseWriter, r *http.Request) {
	//b ,_ :=ioutil.ReadFile("web/web.txt")
	_, _ = fmt.Fprintf(w, "<h1>Hello world</h1>\n"+
		"<button id= 'b1'>click me</button>\n"+
		"<img src='./python_cm.png'>")
	//"<img src='https://upload.wikimedia.org/wikipedia/commons/thumb/a/a5/Basic_concept_of_Kalman_filtering.svg/1920px-Basic_concept_of_Kalman_filtering.svg.png'>")
}
