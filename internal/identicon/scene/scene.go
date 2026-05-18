package scene

import (
	"learning/internal/identicon/generator"
)

type Scene struct {
	Width  int
	Height int

	Background string

	Objects []Object
}

func BuildScene(mask generator.Mask) {

}
