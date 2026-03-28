package gen

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tree struct {
	pos rl.Vector3
	wood rl.Material
	leaves rl.Material
	models []rl.Model
	stemHeight float32
}

func NewTree(pos rl.Vector3) Tree {
	var wood rl.Material = rl.LoadMaterialDefault()
	wood.GetMap(0).Color = rl.Brown

	var leaves rl.Material = rl.LoadMaterialDefault()
	leaves.GetMap(0).Color = rl.NewColor(41, 74, 0,255)

	stemHeight := 4*(1.5+rand.Float32())
	stemWidth := min(float32(1-rand.Float32()*0.75),0.7)

	stem := rl.GenMeshCylinder(stemWidth,stemHeight,8)

	leaf1 := rl.GenMeshSphere(float32(1-(rand.Float32()*0.8))+stemWidth,4+rand.IntN(3),4+rand.IntN(3))
	leaf2 := rl.GenMeshSphere(float32(1+(rand.Float32()*0.5))+stemWidth,4+rand.IntN(3),4+rand.IntN(3))
	leaf3 := rl.GenMeshSphere(float32(1+(rand.Float32()*0.9))+stemWidth,4+rand.IntN(3),4+rand.IntN(3))

	WoodMeshes := []rl.Mesh{stem}
	LeafMeshes := []rl.Mesh{leaf1,leaf2,leaf3}

	models := Compile(WoodMeshes,LeafMeshes,wood,leaves,pos,stemHeight)

	return Tree{pos,wood, leaves, models, stemHeight}
}

func Compile(wmesh, lmesh []rl.Mesh,wood,leaves rl.Material,pos rl.Vector3, stem float32) []rl.Model {
    var models []rl.Model
	for _, obj := range wmesh {
		model := rl.LoadModelFromMesh(obj)
		model.Materials = &wood
		models = append(models, model)
    }
    for i, obj := range lmesh {
		model := rl.LoadModelFromMesh(obj)
		model.Materials = &leaves
		model.Transform = rl.MatrixTranslate(0, pos.Y+stem-(float32(i)*1), 0)
		models = append(models, model)
    }
	return models
}

func (t Tree) Render(debug bool) {
    for _, obj := range t.models {
		if debug {rl.DrawModelWires(obj,t.pos,1,rl.White)} else {rl.DrawModel(obj,t.pos,1,rl.White)}
    }
}