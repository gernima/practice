package primeNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringLength(t *testing.T) {
	a := assert.New(t)

	a.True(IsPrimeNumber(2))
	a.True(IsPrimeNumber(3))
	a.True(IsPrimeNumber(5))
	a.True(IsPrimeNumber(7))
	a.True(IsPrimeNumber(11))
	a.True(IsPrimeNumber(13))
	a.True(IsPrimeNumber(97))

	// НЕ простые числа
	a.False(IsPrimeNumber(4))
	a.False(IsPrimeNumber(6))
	a.False(IsPrimeNumber(9))
	a.False(IsPrimeNumber(15))
	a.False(IsPrimeNumber(100))
}
