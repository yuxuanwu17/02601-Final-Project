package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

// ReadSingleFile Read the single file and transform to [][]float64 pixel and normalize to it in [0-1]
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

// ReadMultipleFiles would receive the fold direction and output the X data and y labels
func ReadMultipleFiles(folderDir string) ([][]float64, [][]float64) {
	files, _ := ioutil.ReadDir(folderDir)
	X := make([][]float64, 0)
	Y := make([][]float64, 0)
	for _, f := range files {
		if f.Name() == ".DS_Store" {
			continue
		}
		char := f.Name()[:1]
		y := OneHotEncoding(char)
		Y = append(Y, y)
		fig := ReadSingleFile(folderDir + "/" + f.Name())
		X = append(X, fig)
	}
	return X, Y
}

// OneHotEncoding would obtain the label from given file name
func OneHotEncoding(char string) []float64 {
	singleLabel := make([]float64, 24)
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
	//fmt.Println(singleLabel)

	return singleLabel
}

// imgToGrey image convert to the grey image and normalize the image by dividing 255
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
	// populate the pixel array and divide by 255 to normalize the data
	for i := 0; i < len(gray.Pix); i++ {
		pixels[i] = float64(gray.Pix[i]) / 255.0
	}
	return pixels
}

// TokenToLabel would accept the index output during prediction and convert to the string format
func TokenToLabel(input int) string {
	dict := make(map[int]string, 24)
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "J", "K",
		"L", "M", "N", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i := 0; i < 24; i++ {
		dict[i] = alphabet[i]
	}
	return dict[input]
}

// ObtainLabelFromString would obtain the actual label from input direction
func ObtainLabelFromString(inputString string) string {
	s := strings.Split(inputString, "/")
	return s[1][:1]
}

// ObtainIndexFromArray this would return the actual label
func ObtainIndexFromArray(input []float64) string {
	for i, i2 := range input {
		if i2 == 1 {
			return TokenToLabel(i)
		}
	}
	return ""
}

// ImageToPNG would transform the image.Gray file to the actual png file
func ImageToPNG(finalImage image.Gray, filename string) {
	outputFile, err := os.Create(filename + ".png")
	if err != nil {
		fmt.Println("Sorry: couldn't create the file!")
		os.Exit(1)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, &finalImage)
	if err != nil {
		fmt.Println(err)
	}

}

// Mtx2Pixel would transform the input pixel matrix
//convert to the image.Gray format and invoke ImageToPNG to output the misclassified file
func Mtx2Pixel(input []float64, label, actualLabel string) {
	byteList := make([]byte, 1152)
	for i := 0; i < len(input); i++ {
		pixels := input[i] * 255
		convert2Byte := byte(pixels)
		byteList[i] = convert2Byte
	}
	var preparedImage image.Gray
	preparedImage.Stride = 24
	rect := image.Rect(0, 0, 24, 48)
	preparedImage.Rect = rect
	preparedImage.Pix = byteList
	ImageToPNG(preparedImage, "plot/actualLabel_"+actualLabel+"_predictedLabel_"+label)
}
