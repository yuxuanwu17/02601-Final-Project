package main

import (
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

func ReadMultipleFiles(folderDir string) ([][]float64, []float64) {
	files, _ := ioutil.ReadDir(folderDir)
	X := make([][]float64, 0)
	Y := make([]float64, 0)
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

func OneHotEncoding(char string) float64 {
	var y float64
	if char == "A" {
		y = 1.0
	}
	if char == "B" {
		y = 2.0
	}
	if char == "C" {
		y = 3.0
	}
	if char == "D" {
		y = 4.0
	}
	if char == "E" {
		y = 5.0
	}
	if char == "F" {
		y = 6.0
	}
	if char == "G" {
		y = 7.0
	}
	if char == "H" {
		y = 8.0
	}
	if char == "J" {
		y = 9.0
	}
	if char == "K" {
		y = 10.0
	}
	if char == "L" {
		y = 11.0
	}
	if char == "M" {
		y = 12.0
	}
	if char == "N" {
		y = 13.0
	}
	if char == "P" {
		y = 14.0
	}
	if char == "Q" {
		y = 15.0
	}
	if char == "R" {
		y = 16.0
	}
	if char == "S" {
		y = 17.0
	}
	if char == "T" {
		y = 18.0
	}
	if char == "U" {
		y = 19.0
	}
	if char == "V" {
		y = 20.0
	}
	if char == "W" {
		y = 21.0
	}
	if char == "X" {
		y = 22.0
	}
	if char == "Y" {
		y = 23.0
	}
	if char == "Z" {
		y = 24.0
	}

	return y
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
