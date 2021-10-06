package main

import (
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	f, err := os.Open("ass2_processed_data/A1.jpeg")
	if err != nil {
		panic(err)
	}

	// decode the figure
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	size := img.Bounds().Size()
	fmt.Println(size)
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
	fmt.Println(len(pixels[0]))
	fmt.Println(len(pixels))
	fmt.Println(pixels)
}
