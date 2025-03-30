package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceNew(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{1, 4, 9}, sliceNew([]int{1, 2, 3}))
	a.Equal([]int{49}, sliceNew([]int{-7}))

	a.Nil(sliceNew(nil))
	a.Equal([]int{}, sliceNew([]int{}))
}

func TestSliceInPlace(t *testing.T) {
	a := assert.New(t)

	slice := []int{1, 2, 3}
	sliceInPlace(slice)
	a.Equal([]int{1, 4, 9}, slice)

	slice = []int{-7}
	sliceInPlace(slice)
	a.Equal([]int{49}, slice)

	var nilSlice []int
	sliceInPlace(nilSlice)
	a.Nil(nilSlice)

	emptySlice := []int{}
	sliceInPlace(emptySlice)
	a.Equal([]int{}, emptySlice)

	original := []int{1, 2, 3}
	result := sliceNew(original)
	a.Equal([]int{1, 2, 3}, original)
	a.Equal([]int{1, 4, 9}, result)
}
