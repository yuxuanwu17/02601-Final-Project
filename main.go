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

	//fmt.Println(img.At(1,1))
	//fmt.Println(reflect.TypeOf(img.At(1,1)))
	//fmt.Println(img.At(1,1).RGBA())
	//fmt.Println(img.At(1,2).RGBA())

	size := img.Bounds().Size()
	fmt.Println(size)
	var pixels [][]int32
	//put pixels into two three two dimensional array
	for i := 0; i < size.X; i++ {
		var y []int32
		for j := 0; j < size.Y; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			//fmt.Println(img.At(i,j))
			//fmt.Println("=====================")
			//fmt.Println(int32(lum/256))
			//
			//fmt.Println("=====================")
			//fmt.Println()
			y = append(y, int32(lum/256))
		}
		pixels = append(pixels, y)
	}

	fmt.Println(pixels)
	//fmt.Println(pixels[0])
	//fmt.Println(reflect.TypeOf(int32(pixels[0])))
}
