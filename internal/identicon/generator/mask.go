package generator

type Mask struct {
	bits []bool
	Size int
}

func (m *Mask) index(x, y int) int {
	return y*m.Size + x
}

func (m *Mask) Filled(x, y int) bool {
	return m.bits[m.index(x, y)]
}

func (m *Mask) Set(x, y int, v bool) {
	m.bits[m.index(x, y)] = v
}

func BuildSymmetricMask(e *Entropy, Size int) Mask {
	mask := Mask{make([]bool, Size*Size), Size}

	halfSize := Size/2 + 1

	for y := 0; y < Size; y++ {
		for x := 0; x < halfSize; x++ {
			isFilled := e.NextBool()
			mask.Set(x, y, isFilled)
			mask.Set(Size-x-1, y, isFilled)
		}
	}
	return mask
}
