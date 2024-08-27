package main

type Ray struct {
	orig *Point3
	dir  *Vec3
}

func (r *Ray) Orig() *Point3 {
	return r.orig
}

func (r *Ray) Dir() *Vec3 {
	return r.dir
}

func (r *Ray) At(t float64) *Point3 {
	return VectorSum(r.orig, (r.dir.MultiplyScalar(t)))
}
