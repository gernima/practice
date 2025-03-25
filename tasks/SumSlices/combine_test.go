package combineSlice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSlices(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{2, 4, 6}, SumSlices([]int{1, 2, 3}, []int{1, 2, 3}))
	a.Equal([]int{11, 22, 33}, SumSlices([]int{10, 20, 30}, []int{1, 2, 3}))
	a.Equal([]int{}, SumSlices([]int{}, []int{}))
	a.Equal([]int{-1, 0, 1}, SumSlices([]int{-1, -1, -1}, []int{0, 1, 2}))
	a.Equal([]int{100}, SumSlices([]int{50}, []int{50}))
	a.Equal([]int{-1, 0, 1, 1}, SumSlices([]int{-1, -1, -1, 1}, []int{0, 1, 2}))
}
