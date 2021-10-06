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

func ReadMultipleFiles(folderDir string) ([][]float64, []string) {
	files, _ := ioutil.ReadDir(folderDir)
	X := make([][]float64, 0)
	Y := make([]string, 0)
	for _, f := range files {
		y := ObtainLabels(f.Name())
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

func ObtainLabels(foldName string) string {
	return foldName[:1]
}

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
