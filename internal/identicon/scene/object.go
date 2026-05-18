package scene

import "io"

type Object interface {
	WriteSVG(io.Writer)
}
