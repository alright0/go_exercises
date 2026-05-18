package renderer

import (
	"fmt"
	"io"
	"learning/internal/identicon/scene"
)

type Renderer struct {
}

func (r Renderer) Render(scene scene.Scene, w io.Writer) error {
	header, err := fmt.Fprintf(w,
		`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d">`,
		scene.Width,
		scene.Height,
	)
	if err != nil {
		return err
	}

	background, err := fmt.Fprintf(w,
		`<rect width="100%%" height="100%%" fill="%s"/>`,
		scene.Background,
	)
	if err != nil {
		return err
	}

	for _, rect := range scene.Rects {
		shape, err := fmt.Fprintf(
			w,
			`<rect x="%f" y="%f" width="%f" height="%f" fill="%s"/>`,
			rect.X,
			rect.Y,
			rect.Width,
			rect.Height,
			rect.Fill,
		)
		if err != nil {
			return err
		}

	}

	bytes, err := io.WriteString(w, "</svg>")
	if err != nil {
		return err
	}

}
