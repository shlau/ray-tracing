package main

import (
	"fmt"
	"io"
	"math"
)

type vec3 struct {
	e []float64
}

// point3 is just an alias for vec3, but useful for geometric clarity in the code.
type point3 = vec3

func NewVec3(e0 float64, e1 float64, e2 float64) *vec3 {
	return &vec3{e: []float64{e0, e1, e2}}
}

func (v *vec3) X() float64 {
	return v.e[0]
}

func (v *vec3) Y() float64 {
	return v.e[1]
}

func (v *vec3) Z() float64 {
	return v.e[2]
}

func (v *vec3) GetAxis(idx int) float64 {
	return v.e[idx]
}

func (v1 *vec3) Negative() *vec3 {
	return &vec3{[]float64{v1.e[0] * -1, v1.e[1] * -1, v1.e[2] * -1}}
}

func (v1 *vec3) Add(v2 *vec3) *vec3 {
	v1.e[0] += v2.e[0]
	v1.e[1] += v2.e[1]
	v1.e[2] += v2.e[2]

	return v1
}

func (v *vec3) MultiplyScalar(val float64) *vec3 {
	v.e[0] *= val
	v.e[1] *= val
	v.e[2] *= val

	return v
}

func (v *vec3) Divide(val float64) *vec3 {
	return v.MultiplyScalar(1 / val)
}

func (v *vec3) lengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v *vec3) Length() float64 {
	return math.Sqrt(v.lengthSquared())
}

// Vector Utility Functions

func WriteVector(w io.Writer, v *vec3) {
	fmt.Fprintf(w, "%f %f %f", v.e[0], v.e[1], v.e[2])
}

func VectorSum(u *vec3, v *vec3) *vec3 {
	return &vec3{[]float64{u.e[0] + v.e[0], u.e[1] + v.e[1], u.e[2] + v.e[2]}}
}

func VectorDiff(u *vec3, v *vec3) *vec3 {
	return &vec3{[]float64{u.e[0] - v.e[0], u.e[1] - v.e[1], u.e[2] - v.e[2]}}
}

func VectorProduct(u *vec3, v *vec3) *vec3 {
	return &vec3{[]float64{u.e[0] * v.e[0], u.e[1] * v.e[1], u.e[2] * v.e[2]}}
}

func VectorScalarProduct(t float64, v *vec3) *vec3 {
	return &vec3{e: []float64{t * v.e[0], t * v.e[1], t * v.e[2]}}
}

func VectorScalarDivision(t float64, v *vec3) *vec3 {
	return VectorScalarProduct(1/t, v)
}

func Dot(u *vec3, v *vec3) float64 {
	return u.e[0]*v.e[0] + u.e[1]*v.e[1] + u.e[2]*v.e[2]
}

func Cross(u *vec3, v *vec3) *vec3 {
	return &vec3{[]float64{u.e[1]*v.e[2] - u.e[2]*v.e[1],
		u.e[2]*v.e[0] - u.e[0]*v.e[2],
		u.e[0]*v.e[1] - u.e[1]*v.e[0]}}
}

func UnitVector(v *vec3) *vec3 {
	return VectorScalarDivision(v.Length(), v)
}
