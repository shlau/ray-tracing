package main

type HittableList struct {
	objects []Hittable
}

func NewHittableList(objs []Hittable) *HittableList {
	return &HittableList{objects: objs}
}

func (h *HittableList) clear() {
	h.objects = nil
}

func (h *HittableList) add(obj Hittable) {
	h.objects = append(h.objects, obj)
}

func (h *HittableList) Hit(r *Ray, rayTMin float64, rayTMax float64, rec *HitRecord) bool {
	tempRec := &HitRecord{}
	hitAnything := false
	closestSoFar := rayTMax

	for _, obj := range h.objects {
		if obj.Hit(r, rayTMin, closestSoFar, tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			rec.p = tempRec.p
			rec.normal = tempRec.normal
			rec.t = tempRec.t
			rec.frontFace = tempRec.frontFace
		}
	}
	return hitAnything
}
