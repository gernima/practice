package squareSlice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquareSliceNew(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{1, 4, 9}, SliceNew([]int{1, 2, 3}))
	a.Equal([]int{0, 1, 100}, SliceNew([]int{0, -1, 10}))
	a.Equal([]int{}, SliceNew([]int{}))
	a.Equal([]int{49}, SliceNew([]int{-7}))
	a.Equal([]int{1, 0, 1}, SliceNew([]int{-1, 0, 1}))
}

func TestSliceInPlace(t *testing.T) {
	a := assert.New(t)

	slice := []int{1, 2, 3}
	SliceInPlace(slice)
	a.Equal([]int{1, 4, 9}, slice)

	slice = []int{0, -1, 10}
	SliceInPlace(slice)
	a.Equal([]int{0, 1, 100}, slice)

	slice = []int{}
	SliceInPlace(slice)
	a.Equal([]int{}, slice)

	slice = []int{-7}
	SliceInPlace(slice)
	a.Equal([]int{49}, slice)

	slice = []int{-1, 0, 1}
	SliceInPlace(slice)
	a.Equal([]int{1, 0, 1}, slice)
}
