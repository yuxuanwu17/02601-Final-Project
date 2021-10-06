package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func ReadSingleFile(fileDir string) [][]float64 {
	f, err := os.Open(fileDir)
	if err != nil {
		panic(err)
	}

	// decode the figure
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	size := img.Bounds().Size()
	var pixels [][]float64
	//put pixels into two three two dimensional array
	for i := 0; i < size.X; i++ {
		var y []float64
		for j := 0; j < size.Y; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			y = append(y, float64(int32(lum/256))/255)
		}
		pixels = append(pixels, y)
	}
	// pixel array is a
	return pixels
}

/*
[][][]float64
*/

func ReadMultipleFiles(folderDir string) {
	files, _ := ioutil.ReadDir(folderDir)
	X := make([][][]float64, 0)
	for _, f := range files {
		fmt.Println(f.Name())
		fig := ReadSingleFile(folderDir + "/" + f.Name())
		X = append(X, fig)
	}
	fmt.Println(len(X))
	fmt.Println(len(X[0]))
	fmt.Println(len(X[0][0]))
}
