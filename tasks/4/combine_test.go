package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSumSlices(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{2, 4, 3}, sumSlices([]int{1, 2, 3}, []int{1, 2}))
	a.Equal([]int{2, 4, 3}, sumSlices([]int{1, 2}, []int{1, 2, 3}))

	a.Equal([]int{}, sumSlices([]int{}, []int{}))
	a.Equal([]int{1, 2}, sumSlices([]int{}, []int{1, 2}))
	a.Equal([]int{1, 2}, sumSlices([]int{1, 2}, []int{}))

	a.Nil(sumSlices(nil, []int{1, 2}))
	a.Nil(sumSlices([]int{1, 2}, nil))
	a.Nil(sumSlices(nil, nil))

	a.Equal([]int{0, 0}, sumSlices([]int{-1, 2}, []int{1, -2}))

	a.Equal([]int{math.MaxInt32}, sumSlices([]int{math.MaxInt32}, []int{0}))

	a.Equal([]int{2, 3, 3}, sumSlices([]int{1, 1}, []int{1, 2, 3}))
	a.Equal([]int{2, 3, 1}, sumSlices([]int{1, 1, 1}, []int{1, 2}))

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	result := sumSlices(s1, s2)
	a.Equal([]int{1, 2, 3}, s1)
	a.Equal([]int{4, 5, 6}, s2)
	a.Equal([]int{5, 7, 9}, result)
}
