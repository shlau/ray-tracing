package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Image
	imageWidth := 256
	imageHeight := 256

	// Render
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := 0; j < imageHeight; j++ {
		log.Printf("\rScanlines remaining: ------------------------------------- %d ", (imageHeight - j))
		for i := 0; i < imageWidth; i++ {
			pixelColor := NewVec3(float64(i)/float64(imageWidth-1), float64(j)/float64(imageHeight-1), 0)
			WriteColor(os.Stdout, pixelColor)
		}
	}

	log.Println("\rDone--------------------------------------------")
}
