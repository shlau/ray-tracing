package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func HitSphere(center *Point3, radius float64, r *Ray) float64 {
	oc := VectorDiff(center, r.Orig())
	a := Dot(r.Dir(), r.Dir())
	b := -2.0 * Dot(r.Dir(), oc)
	c := Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(discriminant)) / (2.0 * a)
	}

}

func RayColor(r *Ray) *Color {
	t := HitSphere(NewVec3(0, 0, -1), 0.5, r)
	if t > 0.0 {
		N := UnitVector(VectorDiff(r.At(t), NewVec3(0, 0, -1)))
		return VectorScalarProduct(0.5, NewVec3(N.X()+1, N.Y()+1, N.Z()+1))
	}

	unitDirection := UnitVector(r.Dir())
	a := 0.5 * (unitDirection.Y() + 1.0)
	u := VectorScalarProduct(1.0-a, NewVec3(1.0, 1.0, 1.0))
	v := VectorScalarProduct(a, NewVec3(0.5, 0.7, 1.0))
	return VectorSum(u, v)
}

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400

	// Calculate the image height, and ensure that it's at least 1.
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	// Camera

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := NewVec3(0, 0, 0)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := NewVec3(viewportWidth, 0, 0)
	viewportV := NewVec3(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := VectorScalarDivision(float64(imageWidth), viewportU)
	pixelDeltaV := VectorScalarDivision(float64(imageHeight), viewportV)

	// Calculate the location of the upper left pixel.
	vpUCenter := VectorScalarDivision(2, viewportU)
	vpVCenter := VectorScalarDivision(2, viewportV)
	viewportUpperLeft := VectorDiff(VectorDiff(VectorDiff(cameraCenter, NewVec3(0, 0, focalLength)), vpUCenter), vpVCenter)
	pixel00Loc := VectorSum(viewportUpperLeft, VectorScalarProduct(0.5, VectorSum(pixelDeltaU, pixelDeltaV)))

	// Render
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := 0; j < imageHeight; j++ {
		log.Printf("\rScanlines remaining: ------------------------------------- %d ", (imageHeight - j))
		for i := 0; i < imageWidth; i++ {
			deltaU := VectorScalarProduct(float64(i), pixelDeltaU)
			deltaV := VectorScalarProduct(float64(j), pixelDeltaV)
			offset := VectorSum(deltaU, deltaV)
			pixelCenter := VectorSum(pixel00Loc, offset)
			rayDirection := VectorSum(pixelCenter, cameraCenter.Negative())
			r := Ray{cameraCenter, rayDirection}
			pixelColor := RayColor(&r)
			WriteColor(os.Stdout, pixelColor)
		}
	}

	log.Println("\rDone--------------------------------------------")
}
