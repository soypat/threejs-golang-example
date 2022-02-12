package main

import three "github.com/soypat/gthree"

type rigidBody struct {
	subBodies []*three.Euler
}

func NewRigidBody(rotations ...*three.Euler) rigidBody {
	return rigidBody{
		subBodies: rotations,
	}
}

func (r *rigidBody) rotate(dx, dy, dz float64) {
	for i := range r.subBodies {
		x := r.subBodies[i].Get("x").Float()
		y := r.subBodies[i].Get("y").Float()
		z := r.subBodies[i].Get("z").Float()
		r.subBodies[i].Set(x+dx, y+dy, z+dz, "")
	}
}
