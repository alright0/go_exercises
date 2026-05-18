package generator

type Generator struct {
	Mask Mask
	Size int
	//GridSize  int
	//ImageSize int
	//Padding   int
}

func (g *Generator) Generate(seed string) Scene {
	entropy := NewEntropy(seed)

	mask := BuildSymmetricMask(entropy, g.Size)

	scene := BuildScene(mask)

	return scene
}
