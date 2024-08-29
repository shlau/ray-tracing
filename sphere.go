package main

import "math"

type Sphere struct {
	center *Point3
	radius float64
}

func NewSphere(c *Point3, r float64) *Sphere {
	return &Sphere{center: c, radius: math.Max(r, 0)}
}

func (s *Sphere) Hit(r *Ray, rayTMin float64, rayTMax float64, rec *HitRecord) bool {
	oc := VectorDiff(s.center, r.Orig())
	a := r.Dir().lengthSquared()
	h := Dot(r.Dir(), oc)
	c := oc.lengthSquared() - s.radius*s.radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)
	// Find the nearest root that lies in the acceptable range.
	root := (h - sqrtd) / a
	if root <= rayTMin || rayTMax <= root {
		root = (h + sqrtd) / a
		if root <= rayTMin || rayTMax <= root {
			return false
		}
	}

	rec.t = root
	rec.p = r.At(rec.t)
	rec.normal = VectorScalarDivision(s.radius, VectorDiff(rec.p, s.center))

	outwardNormal := VectorScalarDivision(s.radius, VectorDiff(rec.p, s.center))
	rec.setFaceNormal(r, outwardNormal)
	return true
}
