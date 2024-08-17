package main

import "fmt"

func main() {
	// Image
	imageWidth := 256
	imageHeight := 256

	// Render
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / (float64(imageWidth - 1))
			g := float64(j) / float64(imageWidth-1)
			b := 0.0

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
