package materia

import (
	three "github.com/soypat/gthree"
)

// color can be specified like so:
// "rgb(255, 0, 0)"
// "rgb(100%, 0%, 0%)"
// "skyblue" // X11 color names (without CamelCase), see three.js source
// "hsl(0, 100%, 50%)"
func ColoredLambertSurface(color string) *three.MeshLambertMaterial {
	boxparam := three.NewMaterialParameters()
	boxparam.Color = three.NewColor(color)
	return three.NewMeshLambertMaterial(boxparam)
}
