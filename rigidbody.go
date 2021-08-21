package main

import "github.com/soypat/three"

type rigidBody struct {
	subBodies []*three.Euler
}

func NewRigidBody(rotations ...*three.Euler) rigidBody {
	return rigidBody{
		subBodies: rotations,
	}
}

func (r *rigidBody) rotate(x, y, z float64) {
	for i := range r.subBodies {
		r.subBodies[i].Set("x", r.subBodies[i].Get("x").Float()+x)
		r.subBodies[i].Set("y", r.subBodies[i].Get("y").Float()+y)
		r.subBodies[i].Set("z", r.subBodies[i].Get("z").Float()+z)
	}
}
