package bvector

import "fmt"

type BVector struct {
	data int
}

// inverse bit in position
func (v *BVector) inverseBitPos(index int) {
	v.data ^= 1 << index
}

func (v *BVector) Show() {
	fmt.Println(fmt.Sprintf("%b", v.data))
}

func (v *BVector) ToNumber() int {
	return v.data
}

func (v *BVector) Length() int {
	i := 0
	data := v.data
	for data != 0 {
		i++
		data >>= 1
	}

	return i
}

// IsSetBit check if bit position is 1
func (v *BVector) IsSetBit(index int) bool {
	return (v.data & (1 << index)) != 0
}

//SetBit set bit position to 1
func (v *BVector) SetBit(index int) {
	v.data |= 1 << index
}

// ResetBit set bit position to 0
func (v *BVector) ResetBit(index int) {
	v.data &= ^(1 << index)
}

// InverseBits inverse all bits
func (v *BVector) InverseBits() {
	for i := 0; i < v.Length(); i++ {
		v.inverseBitPos(i)
	}
}

// RevertBits revert bit in vector
func (v *BVector) RevertBits() int {
	var rev int
	data := v.data
	for data > 0 {
		rev = rev << 1
		if data&1 == 1 {
			rev ^= 1
		}
		data >>= 1
	}

	return rev
}

func (v *BVector) getList() []int {
	data := v.RevertBits()
	length := v.Length()
	var l []int
	for data != 0 {
		l = append(l, data%2)
		data >>= 1
	}

	for length-len(l) > 0 {
		l = append(l, 0)
	}

	return l
}

func main() {
	//v := &BVector{data: 139}
	//fmt.Println(v.getList())
	//v.Show()
	//fmt.Println(fmt.Sprintf("%b", v.RevertBits()))
	//v.Show()
	//fmt.Println(v.ToNumber())
	//fmt.Println(v.IsSetBit(2))
}
