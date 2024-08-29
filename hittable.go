package main

type HitRecord struct {
	p         *Point3
	normal    *Vec3
	t         float64
	frontFace bool
}

type Hittable interface {
	Hit(r *Ray, rayTMin float64, rayTMax float64, rec *HitRecord) bool
}

func (h *HitRecord) setFaceNormal(r *Ray, outwardNormal *Vec3) {
	// Sets the hit record normal vector.
	// NOTE: the parameter `outward_normal` is assumed to have unit length.

	h.frontFace = Dot(r.Dir(), outwardNormal) < 0
	normal := outwardNormal
	if !h.frontFace {
		normal = outwardNormal.Negative()
	}
	h.normal = normal
}
