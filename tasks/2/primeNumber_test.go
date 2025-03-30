package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testIsPrimeNumber(t *testing.T) {
	a := assert.New(t)

	a.True(isPrimeNumber(2))
	a.True(isPrimeNumber(3))
	a.True(isPrimeNumber(5))
	a.True(isPrimeNumber(7))
	a.True(isPrimeNumber(11))
	a.True(isPrimeNumber(13))
	a.True(isPrimeNumber(97))

	// НЕ простые числа
	a.False(isPrimeNumber(4))
	a.False(isPrimeNumber(6))
	a.False(isPrimeNumber(9))
	a.False(isPrimeNumber(15))
	a.False(isPrimeNumber(100))
}
