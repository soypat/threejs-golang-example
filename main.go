package main

import (
	"gjs/materia"

	"github.com/divan/three"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	scale := 200. // Scale will define overall size. All objects will be scaled accordingly.
	logf("scale: %v\n", scale)
	// Get size of window.
	width := js.Global.Get("innerWidth").Float()
	height := js.Global.Get("innerHeight").Float()
	// Bring in the heavyweight renderer.
	// This will render scenes passed into it.
	// You may add it to DOM using appendChild (see end of this program).
	renderer := three.NewWebGLRenderer()
	renderer.SetSize(width, height, true)

	// setup camera and scene
	camera := three.NewPerspectiveCamera(70, width/height, 1, scale*5)
	// put camera along x=y=z line to get nice ISO-view
	camera.Position.Set(scale*5/2, scale*5/2, scale*5/2)
	camera.LookAt(0, 0, 0) // Look at origin so cube is inside view.

	// Scene is passed as an argument to renderer. Scene contains 3D objects.
	scene := three.NewScene()
	// lights, without lights everything will be dark! second and last argument to renderer.
	light := three.NewDirectionalLight(three.NewColor("white"), 1)
	light.Position.Set(0, scale*1.3, scale*1.5)
	scene.Add(light) // This is the idiom to add objects to scene.

	// Create Axis lines
	xline := three.NewLine(lineGeom(zero, xnorm.scale(scale*1.5)), materia.RedLine(1))
	yline := three.NewLine(lineGeom(zero, ynorm.scale(scale*1.5)), materia.GreenLine(1))
	zline := three.NewLine(lineGeom(zero, znorm.scale(scale*1.5)), materia.BlueLine(1))
	scene.Add(xline)
	scene.Add(yline)
	scene.Add(zline)

	// cube object
	geom := three.NewBoxGeometry(&three.BoxGeometryParameters{
		Width:  scale,
		Height: scale,
		Depth:  scale,
	})
	boxmat := materia.ColoredLambertSurface("skyblue")
	mesh := three.NewMesh(geom, boxmat)
	scene.Add(mesh)

	// Generate a rigid body to rotate
	body := NewRigidBody(xline.Rotation, yline.Rotation, zline.Rotation, mesh.Rotation)

	// We create a recursive callback to continuously animate our project.
	var animate func()
	animate = func() {
		body.rotate(0.01, 0.01, 0)
		renderer.Render(scene, camera)
		js.Global.Call("requestAnimationFrame", animate)
	}
	// Add renderer to DOM (HTML page).
	js.Global.Get("document").Get("body").Call("appendChild", renderer.Get("domElement"))
	// start animation using recursive callback method.
	// Each time a frame is finished rendering a new request to animate is called.
	animate()

}

// Line geometry creation.

type vec struct {
	x, y, z float64
}

func (v *vec) scale(a float64) *vec {
	v.x *= a
	v.y *= a
	v.z *= a
	return v
}

var (
	zero  = &vec{}
	xnorm = &vec{1, 0, 0}
	ynorm = &vec{0, 1, 0}
	znorm = &vec{0, 0, 1}
)

func lineGeom(to, from *vec) *three.BasicGeometry {
	geom := three.NewBasicGeometry(three.BasicGeometryParams{})
	geom.AddVertice(to.x, to.y, to.z)
	geom.AddVertice(from.x, from.y, from.z)
	return &geom
}
