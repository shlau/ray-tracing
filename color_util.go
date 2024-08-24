package main

import (
	"fmt"
	"io"
)

type Color = Vec3

func WriteColor(w io.Writer, pixelColor *Color) {
	r := pixelColor.X()
	g := pixelColor.Y()
	b := pixelColor.Z()

	// Translate the [0,1] component values to the byte range [0,255].
	rbyte := int(255.999 * r)
	gbyte := int(255.999 * g)
	bbyte := int(255.999 * b)

	fmt.Fprintf(w, "%d %d %d\n", rbyte, gbyte, bbyte)
}
