package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
)

/*
Read the single file and transform to [][]float64 pixel and normalize to it in [0-1]
*/

func ReadSingleFile(fileDir string) []float64 {
	f, err := os.Open(fileDir)
	if err != nil {
		panic(err)
	}

	// decode the figure
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	pixels := imgToGrey(img)
	return pixels
}

/*
Input: folder direction
Output: X [][][]float64, Y []string
*/

func ReadMultipleFiles(folderDir string) ([][]float64, [][]int) {
	files, _ := ioutil.ReadDir(folderDir)
	X := make([][]float64, 0)
	Y := make([][]int, 0)
	for _, f := range files {
		char := f.Name()[:1]
		y := OneHotEncoding(char)
		Y = append(Y, y)
		fig := ReadSingleFile(folderDir + "/" + f.Name())
		X = append(X, fig)
	}
	//fmt.Println(len(X))
	//fmt.Println(len(X[0]))
	//fmt.Println(len(X[0][0]))
	//fmt.Println(Y)
	//fmt.Println(X)
	return X, Y
}

/*
obtain the label from given file name
*/

func OneHotEncoding(char string) []int {
	singleLabel := make([]int, 24)
	if char == "A" {
		singleLabel[0] = 1
	}
	if char == "B" {
		singleLabel[1] = 1
	}
	if char == "C" {
		singleLabel[2] = 1
	}
	if char == "D" {
		singleLabel[3] = 1
	}
	if char == "E" {
		singleLabel[4] = 1
	}
	if char == "F" {
		singleLabel[5] = 1
	}
	if char == "G" {
		singleLabel[6] = 1
	}
	if char == "H" {
		singleLabel[7] = 1
	}
	if char == "J" {
		singleLabel[8] = 1
	}
	if char == "K" {
		singleLabel[9] = 1
	}
	if char == "L" {
		singleLabel[10] = 1
	}
	if char == "M" {
		singleLabel[11] = 1
	}
	if char == "N" {
		singleLabel[12] = 1
	}
	if char == "P" {
		singleLabel[13] = 1
	}
	if char == "Q" {
		singleLabel[14] = 1
	}
	if char == "R" {
		singleLabel[15] = 1
	}
	if char == "S" {
		singleLabel[16] = 1
	}
	if char == "T" {
		singleLabel[17] = 1
	}
	if char == "U" {
		singleLabel[18] = 1
	}
	if char == "V" {
		singleLabel[19] = 1
	}
	if char == "W" {
		singleLabel[20] = 1
	}
	if char == "X" {
		singleLabel[21] = 1
	}
	if char == "Y" {
		singleLabel[22] = 1
	}
	if char == "Z" {
		singleLabel[23] = 1
	}
	fmt.Println(singleLabel)

	return singleLabel
}

/*
image convert to the grey image and normalize the image by dividing 255
*/

func imgToGrey(img image.Image) []float64 {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	// make a pixel array
	pixels := make([]float64, len(gray.Pix))
	// populate the pixel array and divide by 255 to normalize the data (pixels[i] = (float64(255-gray.Pix[i]) / 255.0 * 0.99) + 0.01)
	for i := 0; i < len(gray.Pix); i++ {
		pixels[i] = float64(gray.Pix[i]) / 255.0
	}
	return pixels
}
