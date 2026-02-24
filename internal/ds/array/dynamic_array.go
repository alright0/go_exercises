package ds

type DynamicArray struct {
	data []int
}

func (d *DynamicArray) Length() int {
	return len(d.data)
}

func (d *DynamicArray) Push(value int) {
	arrSize := len(d.data)
	if arrSize == cap(d.data) {
		arrCap := max(cap(d.data)*2, 1)
		newArr := make([]int, arrSize, arrCap)
		copy(newArr, d.data)
		d.data = newArr
	}

	d.data = d.data[:arrSize+1]
	d.data[arrSize] = value
}

func (d *DynamicArray) Peek() int {
	arrSize := len(d.data)
	return d.data[arrSize-1]
}

func (d *DynamicArray) Pop() int {
	arrSize := len(d.data)

	newSize := arrSize - 1
	value := d.data[newSize]
	d.data = d.data[:newSize]
	if cap(d.data) > 0 && newSize < cap(d.data)/4 {
		arrCap := max(cap(d.data)/2, 1)
		newArr := make([]int, newSize, arrCap)
		copy(newArr, d.data)
		d.data = newArr
	}
	return value
}
